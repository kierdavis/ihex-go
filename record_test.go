package ihex

import (
	"testing"
)

func TestDecodeRecord(t *testing.T) {
	input := ":1004F0003E010894611C711CD501C4010197A1093A"
	rec, err := DecodeRecordHex(input)
	if err != nil {
		t.Error("DecodeRecordHex returned an error: %s", err.Error())
	}

	if rec.Type != 0x00 {
		t.Error("rec.Type: expected 0x00, got 0x%02X", rec.Type)
	}

	if rec.Address != 0x04F0 {
		t.Error("rec.Address: expected 0x04F0, got 0x%02X", rec.Address)
	}

	if len(rec.Data) != 16 {
		t.Error("rec.Data: expected length 16, got length %d", len(rec.Data))
	}

	if rec.Data[0] != 0x3E {
		t.Error("rec.Data[0]: expected 0x3E, got 0x%02X", rec.Data[0])
	}

	if rec.Data[1] != 0x01 {
		t.Error("rec.Data[1]: expected 0x01, got 0x%02X", rec.Data[1])
	}

	if rec.Data[15] != 0x09 {
		t.Error("rec.Data[15]: expected 0x09, got 0x%02X", rec.Data[15])
	}
}
