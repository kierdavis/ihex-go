package ihex

func Checksum(buf []byte) (sum byte) {
    for _, x := range buf {
        sum += x
    }
    
    return -sum
}
