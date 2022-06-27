package util

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
)

const (
	CONST_UINT64_BYTE_NUM = 8
)

func Recv(reader io.Reader, maxSize uint64) (data []byte, err error) {
	size, err := ReadSize(reader)
	if err != nil {
		return nil, err
	}
	if maxSize > 0 && size > maxSize {
		return nil, fmt.Errorf("rpc: data size overflows maxSize(%d)", maxSize)
	}
	if size == 0 {
		return
	}
	data = make([]byte, size)
	if err = Read(reader, data); err != nil {
		return nil, err
	}
	return data, nil
}

func Send(writer io.Writer, data []byte) (err error) {
	var size [binary.MaxVarintLen64]byte
	if data == nil || len(data) == 0 {
		n := binary.PutUvarint(size[:], uint64(0))
		return Write(writer, size[:n], false)
	}

	n := binary.PutUvarint(size[:], uint64(len(data)))
	err = Write(writer, size[:n], false)
	if err != nil {
		return err
	}
	return Write(writer, data, false)
}

func ReadSize(reader io.Reader) (uint64, error) {
	var (
		size  uint64
		shift uint64
	)
	for i := 1; ; i++ {
		data, err := ReadByte(reader)
		if err != nil {
			return 0, err
		}
		if data < 0x80 {
			if i == CONST_UINT64_BYTE_NUM && data&0x80 > 0 {
				return 0, errors.New("rpc: header size overflows a 64-bit integer")
			}
			return size | uint64(data)<<shift, nil
		}
		
		size |= uint64(data&0x7F) << shift
		shift += 7
	}
}

func ReadByte(reader io.Reader) (byte, error) {
	buff := make([]byte, 1)
	if err := Read(reader, buff); err != nil {
		return 0, err
	}
	return buff[0], nil
}

func Read(reader io.Reader, buff []byte) error {
	for i := 0; i < len(buff); {
		n, err := reader.Read(buff[i:])
		if err != nil {
			if nerr, ok := err.(net.Error); !ok || !nerr.Temporary() {
				return err
			}
		}
		i += n
	}
	return nil
}

func Write(writer io.Writer, data []byte, onePacket bool) error {
	if onePacket {
		if _, err := writer.Write(data); err != nil {
			return nil
		}
		return nil
	}
	for i := 0; i < len(data); {
		n, err := writer.Write(data[i:])
		if err != nil {
			if nerr, ok := err.(net.Error); !ok || !nerr.Temporary() {
				return err
			}
		}
		i += n
	}
	return nil
}

func MaxUint32(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}