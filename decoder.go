package ihex

import (
    "bufio"
    "fmt"
    "io"
    "strings"
)

// Parses IHEX files.
type Decoder struct {
    scanner *bufio.Scanner
    rec Record
    err error
    lineno int
}

// Create a new Decoder using input from the given io.Reader.
func NewDecoder(r io.Reader) (d *Decoder) {
    scanner := bufio.NewScanner(r)
    return &Decoder{
        scanner: scanner,
        lineno: 0,
    }
}

// Read and decode one record from the source, returning true if the record was
// decoded successfully. False may be returned if the end of the file is reached
// or an error has occurred. This method is designed to be called repeatedly
// to read the entire file, and then have the Err method called to determine
// if an error occurred. The scanned record can be retrieved with the Record
// method.
func (d *Decoder) Scan() (ok bool) {
    if d.err != nil || !d.scanner.Scan() {
        return false
    }
    
    d.lineno++
    
    line := strings.Trim(d.scanner.Text(), " \t\r\n")
    if len(line) == 0 {
        // try again
        return d.Scan()
    }
    
    if line[0] != ':' {
        d.err = fmt.Errorf("parse error at line %d: line does not begin with a colon", d.lineno)
        return false
    }
    
    rec, err := DecodeRecordHex(line)
    if err != nil {
        d.err = fmt.Errorf("parse error at line %d: %s", d.lineno, err.Error())
        return false
    } else {
        d.rec = rec
        return true
    }
    
}

// Returns the most recently parsed record. An empty record is returned if Scan
// has not yet been called, or if Scan has never returned true.
func (d *Decoder) Record() (rec Record) {
    return d.rec
}

// Returns the first error that occurred during parsing, or nil if there was no
// error.
func (d *Decoder) Err() (err error) {
    if d.err == nil {
        return d.scanner.Err()
    }
    return d.err
}
