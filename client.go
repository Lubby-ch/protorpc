package protorpc

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"protorpc/wire"

	"net"
	"net/rpc"
	"sync"
)

type clientCodec struct {
	conn io.ReadWriteCloser

	respHeader *wire.ResponseHeader
	// Protobuf-RPC responses include the request id but not the request method.
	// Package rpc expects both.
	// We save the request method in pending when sending a request
	// and then look it up by request ID when filling out the rpc Response.
	mutex   sync.Mutex // protects seq, pending
	pending map[uint64]string
}

func NewClientCodec(conn io.ReadWriteCloser) rpc.ClientCodec {
	return &clientCodec{
		conn:    conn,
		pending: make(map[uint64]string),
	}
}

// WriteRequest handles user's request and send request.Seq and request.ServiceMethod to the remote server to realize
// Remote procedure call.
//@param request
func (c clientCodec) WriteRequest(request *rpc.Request, i interface{}) error {
	c.mutex.Lock()
	c.pending[request.Seq] = request.ServiceMethod
	c.mutex.Unlock()

	var (
		req proto.Message
	)
	if i != nil {
		var ok bool
		req, ok = i.(proto.Message)
		if !ok {
			return fmt.Errorf("rpc.ClientCodec.WriteRequest: %T does not implement proto.Message", i)
		}
	}
	return writeRequest(c.conn, request.Seq, request.ServiceMethod, req)
}

func (c clientCodec) ReadResponseHeader(response *rpc.Response) error {
	var (
		header = new(wire.ResponseHeader)
	)
	err := readResponseHeader(c.conn, header)
	if err != nil {
		return nil
	}

	c.mutex.Lock()
	response.Seq = header.Id
	response.Error = header.Error
	response.ServiceMethod = c.pending[response.Seq] // when can not find Service Method, "" is permitted
	delete(c.pending, response.Seq)
	c.mutex.Unlock()

	c.respHeader = header
	return nil
}

func (c clientCodec) ReadResponseBody(i interface{}) error {
	var response proto.Message
	if i != nil {
		var ok bool
		response, ok = i.(proto.Message)
		if !ok {
			return fmt.Errorf("rpc.ServerCodec.ReadResponseBody: %T does not implement proto.Message", i)
		}
	}

	err := readResponseBody(c.conn, c.respHeader, response)
	if err != nil {
		return err
	}

	c.respHeader.Reset()
	return nil
}

func (c clientCodec) Close() error {
	return c.conn.Close()
}

// NewClient returns a new rpc.Client to handle requests to the
// set of services at the other end of the connection.
func NewClient(conn io.ReadWriteCloser) *rpc.Client {
	return rpc.NewClientWithCodec(NewClientCodec(conn))
}

// Dial connects to a protobuf-RPC server at the specified network address.
func Dial(network, address string) (*rpc.Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return NewClient(conn), err
}
