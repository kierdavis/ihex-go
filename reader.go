package ihex

import (
	"bufio"
	"io"
)

type Reader struct {
	br *bufio.Reader
}

func NewReader(r io.Reader) (rr *Reader) {
	br := bufio.NewReader(r)
	return &Reader{
		br: br,
	}
}

func (r *Reader) ReadRecord() (rec Record) {
	
}
