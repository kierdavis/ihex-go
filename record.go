package ihex

import (
	"fmt"
)

var ChecksumMismatchError = fmt.Errorf("DecodeRecord: checksum mismatch")

type RecordType uint8

const (
	Data RecordType = 0x00
	EndOfFile RecordType = 0x01
)

type Record struct {
	Address uint16
	Type RecordType
	Data []byte
}

func DecodeRecord(bytes []byte) (rec Record, err error) {
	dataLen := int(bytes[0])
	address := (uint16(bytes[1]) << 8) | uint16(bytes[2])
	recType := RecordType(bytes[3])
	
	checksumIndex := 4 + dataLen
	if bytes[checksumIndex] != Checksum(bytes[:checksumIndex]) {
		return rec, ChecksumMismatchError
	}
	
	data := make([]byte, dataLen)
	copy(data, bytes[4:checksumIndex])
	
	return Record{
		Address: address,
		Type: recType,
		Data: data,
	}, nil
}

func (rec Record) Encode() (bytes []byte) {
	bytes = make([]byte, len(rec.Data) + 5)
	bytes[0] = byte(len(rec.Data))
	bytes[1] = byte(rec.Address >> 8)
	bytes[2] = byte(rec.Address)
	bytes[3] = byte(rec.Type)
	copy(bytes[4:], rec.Data)
	
	last := len(bytes) - 1
	bytes[last] = Checksum(bytes[:last])
	return bytes
}
