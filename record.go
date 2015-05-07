package ihex

import (
	"encoding/hex"
	"fmt"
	"strings"
)

//RecordType is a value from 0 to 5 as described in
//the intel hex spec (http://www.piclist.com/techref/fileext/hex/intel.htm)
type RecordType uint8

const (
	Data                   RecordType = 0x00
	EndOfFile                         = 0x01
	ExtendedSegmentAddress            = 0x02
	StartSegmentAddress               = 0x03
	ExtendedLinearAddress             = 0x04
	StartLinearAddress                = 0x05
)

//ExtAddress represents an extended address of width up to 32-bits
type ExtAddress uint32

//ExtAddressSlice is defined in order to allow implemention of the Sort
//interface for client convenience. Invoke as follows:
//  import "sort"
//	addressesToSort := []ExtAddress{....}
//	sort.Sort(ExtAddressSlice(addressesToSort))
type ExtAddressSlice []ExtAddress
func (a ExtAddressSlice) Len() int           { return len(a) }
func (a ExtAddressSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ExtAddressSlice) Less(i, j int) bool { return a[i] < a[j] }

// Represents a single IHEX record.
type Record struct {
	// The type of record
	Type RecordType

	// The address associated with the record's data (for Data records).
	Address uint16

	// The Extended address associated with the record's data
	ExtendedAddress ExtAddress

	// The payload of the record.
	Data []byte
}

// Interpreting bytes as the binary form of an IHEX record, decode and return a
// Record. err may be non-nil if there was a checksum mismatch.
func DecodeRecord(bytes []byte) (rec Record, err error) {
	dataLen := int(bytes[0])
	address := (uint16(bytes[1]) << 8) | uint16(bytes[2])
	recType := RecordType(bytes[3])

	checksumIndex := 4 + dataLen
	if bytes[checksumIndex] != Checksum(bytes[:checksumIndex]) {
		return rec, fmt.Errorf("checksum mismatch")
	}

	data := make([]byte, dataLen)
	copy(data, bytes[4:checksumIndex])

	return Record{
		Type:    recType,
		Address: address,
		Data:    data,
	}, nil
}

// Interpreting hexStr as the hexadecimal form of an IHEX record with or without
// the leading start token (colon), deocde and return a Record. err may be
// non-nil if there was a checksum mismatch or if the decoding from hexadecimal
// failed.
func DecodeRecordHex(hexStr string) (rec Record, err error) {
	bytes, err := hex.DecodeString(strings.TrimLeft(hexStr, ":"))
	if err != nil {
		return rec, err
	}
	return DecodeRecord(bytes)
}

// Encode the record to the binary form of an IHEX record.
func (rec Record) Encode() (bytes []byte) {
	bytes = make([]byte, len(rec.Data)+5)
	bytes[0] = byte(len(rec.Data))
	bytes[1] = byte(rec.Address >> 8)
	bytes[2] = byte(rec.Address)
	bytes[3] = byte(rec.Type)
	copy(bytes[4:], rec.Data)

	checksumIndex := 4 + len(rec.Data)
	bytes[checksumIndex] = Checksum(bytes[:checksumIndex])
	return bytes
}

// Encode the record to the hexadecimal form of an IHEX record, with the leading
// start token (colon).
func (rec Record) EncodeHex() (hexStr string) {
	return ":" + hex.EncodeToString(rec.Encode())
}

func (rec *Record) String() string {
	return fmt.Sprintf("ExtendedAddress: %x / Address: %x / Data %x", rec.ExtendedAddress, rec.Address, rec.Data)
}
