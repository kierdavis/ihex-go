package ihex

import (
    "os"
)

func ExampleDecoder() {
    d := NewDecoder(os.Stdin)
    
    for d.Scan() {
        // do something with d.Record()
    }
    
    // check if there was an error during parsing
    err := d.Err()
    if err != nil {
        // handle the error
    }
}
