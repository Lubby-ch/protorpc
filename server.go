package protorpc

import (
	"errors"
	"fmt"
	"github.com/Lubby-ch/protorpc/wire"
	"github.com/golang/protobuf/proto"
	"io"
	"net/rpc"
	"sync"
)

type serverCodec struct {
	conn io.ReadWriteCloser

	reqHeader *wire.RequestHeader
	// Package rpc expects uint64 request IDs.
	// We assign uint64 sequence numbers to incoming requests
	// but save the original request ID in the pending map.
	// When rpc responds, we use the sequence number in
	// the response to find the original request ID.
	mutex   sync.Mutex // protects seq, pending
	seq     uint64
	pending map[uint64]uint64
}

// NewServerCodec returns a serverCodec that communicates with the ClientCodec
// on the other end of the given conn.
func NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec {
	return &serverCodec{
		conn:    conn,
		pending: make(map[uint64]uint64),
	}
}

func (s *serverCodec) ReadRequestHeader(req *rpc.Request) error {
	header := new(wire.RequestHeader)
	err := readRequestHeader(s.conn, header)
	if err != nil {
		return err
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.seq++
	s.pending[s.seq] = header.Id
	s.reqHeader = header

	req.ServiceMethod = s.reqHeader.Method
	req.Seq = s.seq
	return nil
}

func (s *serverCodec) ReadRequestBody(i interface{}) error {
	if i == nil {
		return nil
	}
	body, ok := i.(proto.Message)
	if !ok {
		return fmt.Errorf("rpc.ServerCodec.ReadRequestBody: %T does not implement proto.Message", i)
	}
	err := readRequestBody(s.conn, s.reqHeader, body)
	if err != nil {
		return nil
	}
	s.reqHeader.Reset()
	return nil
}

func (s *serverCodec) WriteResponse(resp *rpc.Response, i interface{}) error {
	var response proto.Message
	if i != nil {
		var ok bool
		response, ok = i.(proto.Message)
		if !ok {
			if _, ok = i.(struct{}); !ok {
				s.mutex.Lock()
				delete(s.pending, resp.Seq)
				s.mutex.Unlock()
			}
			return fmt.Errorf("rpc.ServerCodec.WriteResponse: %T does not implement proto.Message", i)
		}
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	id, ok := s.pending[resp.Seq]
	if !ok {
		return fmt.Errorf("rpc: invalid sequence number in response")
	}
	err := writeResponse(s.conn, id, resp.Error, response)
	if err != nil {
		return err
	}
	return nil
}

func (s *serverCodec) Close() error {
	return s.conn.Close()
}

// ServeConn runs the Protocol-RPC server on a single connection.
// ServeConn blocks, serving the connection until the client hangs up.
// The caller typically invokes ServeConn in a go statement.
func ServeConn(conn io.ReadWriteCloser) {
	rpc.ServeCodec(NewServerCodec(conn))
}