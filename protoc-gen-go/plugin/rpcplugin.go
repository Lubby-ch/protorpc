package plugin

import (
	"google.golang.org/protobuf/compiler/protogen"
)

const (
	rpcPackage = protogen.GoImportPath("net/rpc")
)

type rpcPlugin struct {
	*protogen.Plugin
	file  *protogen.File
	gfile *protogen.GeneratedFile
}

func NewRpcPlugin(gen *protogen.Plugin, file *protogen.File, gfile *protogen.GeneratedFile) *rpcPlugin {
	return &rpcPlugin{
		Plugin: gen,
		file:   file,
		gfile:  gfile, // Empty is not allowed
	}
}

func (plugin *rpcPlugin) GenServiceCode() {
	if plugin.gfile == nil { // Empty is not allowed
		return
	}
	if len(plugin.file.Services) == 0 {
		return
	}
	plugin.gfile.P("// Reference imports to suppress errors if they are not otherwise used.")
	plugin.gfile.P("var _ ", rpcPackage.Ident("Server"))
	plugin.gfile.P()

	plugin.gfile.P("// This is a compile-time assertion to ensure that this generated file")
	plugin.gfile.P("// is compatible with the grpc package it is being compiled against.")
	plugin.gfile.P()
	for _, service := range plugin.file.Services {
		plugin.genServiceInterface(service)
		plugin.genServiceServer(service)
		plugin.genServiceClient(service)
	}
}

func (plugin *rpcPlugin) genServiceInterface(service *protogen.Service) {
	serviceName := service.GoName + "Service"
	plugin.gfile.P("type ", serviceName, " interface {")
	for _, method := range service.Methods {
		plugin.gfile.Annotate(serviceName+"."+method.GoName, method.Location)
		plugin.gfile.P(method.GoName, " (in *", plugin.gfile.QualifiedGoIdent(method.Input.GoIdent), ", ", "out *",
			plugin.gfile.QualifiedGoIdent(method.Output.GoIdent), ") ", "error")
	}
	plugin.gfile.P("}")
	plugin.gfile.P()
}

func (plugin *rpcPlugin) genServiceServer(service *protogen.Service) {
	interfaceName := service.GoName + "Service"
	serviceName := "Default" + interfaceName
	plugin.gfile.P("type ", serviceName, " struct {}")
	plugin.gfile.P()

	for _, method := range service.Methods {
		plugin.gfile.P("func (d *", serviceName, ") ", method.GoName, " ( in *", plugin.gfile.QualifiedGoIdent(method.Input.GoIdent),
			", ", "out *", plugin.gfile.QualifiedGoIdent(method.Output.GoIdent), ") ", "error {")
		plugin.gfile.P("panic(\"", serviceName, ".", method.GoName, " is not implemented\")")
		plugin.gfile.P("}")
		plugin.gfile.P()
	}

	plugin.gfile.P("// Register", serviceName, "publish the given ", service.GoName, " implementation on the server.")
	plugin.gfile.P("func Register", serviceName, " (server *rpc.Server, i ", interfaceName, ") error {")
	plugin.gfile.P("if err := server.RegisterName(\"", service.GoName, "\", i); err != nil {")
	plugin.gfile.P("return err")
	plugin.gfile.P("}")
	plugin.gfile.P("return nil")
	plugin.gfile.P("}")
	plugin.gfile.P()
}

func (plugin *rpcPlugin) genServiceClient(service *protogen.Service) {
	plugin.gfile.P("type ", service.GoName, "Client struct {")
	plugin.gfile.P("*rpc.Client")
	plugin.gfile.P("}")
	plugin.gfile.P()

	for _, method := range service.Methods {
		plugin.gfile.P("func (c *", service.GoName, "Client) ", method.GoName, "(in *", plugin.gfile.QualifiedGoIdent(method.Input.GoIdent),
			") (out *", plugin.gfile.QualifiedGoIdent(method.Output.GoIdent), ", err error) {")
		plugin.gfile.P("if in == nil {")
		plugin.gfile.P("in = new(", plugin.gfile.QualifiedGoIdent(method.Input.GoIdent), ")")
		plugin.gfile.P("}")
		plugin.gfile.P()
		plugin.gfile.P("out = new(", plugin.gfile.QualifiedGoIdent(method.Output.GoIdent), ")")
		plugin.gfile.P("if err = c.Call(\"", service.GoName, ".", method.GoName, "\", in, out); err != nil {")
		plugin.gfile.P("return nil, err")
		plugin.gfile.P("}")
		plugin.gfile.P("return out, nil")
		plugin.gfile.P("}")
		plugin.gfile.P()
	}
}
