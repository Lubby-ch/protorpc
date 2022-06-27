package main

import (
	"flag"
	"fmt"
	gengo "google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
	"google.golang.org/protobuf/compiler/protogen"
	"protoc-gen-go/internal_gengo"
	"protoc-gen-go/plugin"
	"strings"
)

type CodePlugin interface {
	GenServiceCode()
}

func GenerateServiceFile(plugin CodePlugin) {
	plugin.GenServiceCode()
}

func main() {
	var (
		flags   flag.FlagSet
		plugins = flags.String("plugins", "", "deprecated option")
	)
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		var (
			gfile *protogen.GeneratedFile
			pbrpc bool
		)
		for _, p := range strings.Split(*plugins, ",") {
			switch p {
			case "plugin":
				pbrpc = true
			case "":
			default:
				return fmt.Errorf("protoc-gen-go: unknown plugin %q", p)
			}
		}
		for _, f := range gen.Files {
			if f.Generate {
				gfile = internal_gengo.GenerateFile(gen, f)
				if pbrpc && gfile != nil {
					GenerateServiceFile(plugin.NewRpcPlugin(gen, f, gfile))
				}
			}
		}
		gen.SupportedFeatures = gengo.SupportedFeatures
		return nil
	})
}

//  protoc --protorpc_out=plugins=plugin:. .\*.proto