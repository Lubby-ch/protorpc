package protorpc

import (
	"fmt"
	"github.com/Lubby-ch/protorpc/util"
	"github.com/Lubby-ch/protorpc/wire"
	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"
	"hash/crc32"
	"io"
)

var (
	UseSnappy            = true
	UseCrc32ChecksumIEEE = true
)

const (
	CONST_REQUEST_HEADER_MAX_LEN = 1024
)

func writeRequest(writer io.Writer, id uint64, method string, request proto.Message) (err error) {
	var (
		protoReq         = []byte{}
		comoressProtoReq = []byte{}
	)

	if request != nil {
		protoReq, err = proto.Marshal(request)
		if err != nil {
			return err
		}
	}

	comoressProtoReq = snappy.Encode(nil, protoReq)

	header := &wire.RequestHeader{
		Id:                         id,
		Method:                     method,
		RawRequestLen:              uint32(len(protoReq)),
		SnappyCompressedRequestLen: uint32(len(comoressProtoReq)),
		Checksum:                   crc32.ChecksumIEEE(comoressProtoReq),
	}

	if !UseSnappy || header.RawRequestLen < header.SnappyCompressedRequestLen {
		header.SnappyCompressedRequestLen = 0
		comoressProtoReq = protoReq
	}
	if !UseCrc32ChecksumIEEE {
		header.Checksum = 0
	}

	protoHeader, err := proto.Marshal(header)
	if err != nil {
		return err
	}

	if len(protoHeader) > CONST_REQUEST_HEADER_MAX_LEN {
		return fmt.Errorf("rpc.writeRequest: the header length: %d is larger than the limit of header : %d.", len(protoHeader), CONST_REQUEST_HEADER_MAX_LEN)
	}

	if err = util.Send(writer, protoHeader); err != nil {
		return nil
	}
	return util.Send(writer, comoressProtoReq)
}

func readRequestHeader(reader io.Reader, header proto.Message) (err error) {
	protoHeader, err := util.Recv(reader, 1024)
	if err != nil {
		return err
	}
	return proto.Unmarshal(protoHeader, header)
}

func readRequestBody(reader io.Reader, header *wire.RequestHeader, body proto.Message) (err error) {
	maxBodyLen := util.MaxUint32(header.RawRequestLen, header.SnappyCompressedRequestLen)

	compressProtoBody, err := util.Recv(reader, uint64(maxBodyLen))
	if err != nil {
		return err
	}

	if header.Checksum != 0 {
		if crc32.ChecksumIEEE(compressProtoBody) != header.Checksum {
			return fmt.Errorf("rpc.readRequestBody: checksum err ")
		}
	}

	var protoBody = compressProtoBody
	if header.SnappyCompressedRequestLen != 0 {
		// 解压
		protoBody, err = snappy.Decode(nil, compressProtoBody)
		if err != nil {
			return err
		}
		// check wire header: rawMsgLen
		if uint32(len(protoBody)) != header.RawRequestLen {
			return fmt.Errorf("rpc.readRequestBody: Unexcpeted header.RawResponseLen.")
		}
	}

	if body != nil {
		err = proto.Unmarshal(protoBody, body)
		if err != nil {
			return nil
		}
	}
	return nil
}

func writeResponse(writer io.Writer, id uint64, strErr string, response proto.Message) (err error) {
	if strErr != "" {
		response = nil
	}

	var (
		protoResp         = []byte{}
		compressProtoResp = []byte{}
	)
	if response != nil {
		protoResp, err = proto.Marshal(response)
		if err != nil {
			return err
		}
	}

	compressProtoResp = snappy.Encode(nil, protoResp)

	header := &wire.ResponseHeader{
		Id:                          id,
		Error:                       strErr,
		RawResponseLen:              uint32(len(protoResp)),
		SnappyCompressedResponseLen: uint32(len(compressProtoResp)),
		Checksum:                    crc32.ChecksumIEEE(compressProtoResp),
	}

	if !UseSnappy {
		header.SnappyCompressedResponseLen = 0
		compressProtoResp = protoResp
	}
	if !UseCrc32ChecksumIEEE {
		header.Checksum = 0
	}

	protoHeader, err := proto.Marshal(header)
	if err != nil {
		return nil
	}
	if len(protoHeader) > CONST_REQUEST_HEADER_MAX_LEN {
		return fmt.Errorf("rpc.writeResponse: the header length: %d is larger than the limit of header : %d.", len(protoHeader), CONST_REQUEST_HEADER_MAX_LEN)
	}
	if err = util.Send(writer, protoHeader); err != nil {
		return nil
	}

	return util.Send(writer, compressProtoResp)
}

func readResponseHeader(reader io.Reader, header proto.Message) (err error) {
	protoHeader, err := util.Recv(reader, 1024)
	if err != nil {
		return err
	}
	return proto.Unmarshal(protoHeader, header)
}

func readResponseBody(reader io.Reader, header *wire.ResponseHeader, body proto.Message) (err error) {
	maxBodyLen := util.MaxUint32(header.RawResponseLen, header.SnappyCompressedResponseLen)

	compressProtoBody, err := util.Recv(reader, uint64(maxBodyLen))
	if err != nil {
		return err
	}

	if header.Checksum != 0 {
		if crc32.ChecksumIEEE(compressProtoBody) != header.Checksum {
			return fmt.Errorf("rpc.readRequestBody: checksum err ")
		}
	}

	var protoBody = compressProtoBody
	if header.SnappyCompressedResponseLen != 0 {
		protoBody, err = snappy.Decode(nil, compressProtoBody)
		if err != nil {
			return err
		}
		// check wire header: rawMsgLen
		if uint32(len(protoBody)) != header.RawResponseLen {
			return fmt.Errorf("rpc.readResponseBody: Unexcpeted header.RawResponseLen.")
		}
	}

	if body != nil {
		err = proto.Unmarshal(protoBody, body)
		if err != nil {
			return nil
		}
	}
	return nil
}
