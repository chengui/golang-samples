package popcount

var pc [256]byte

func init() {
    for i := range pc {
        pc[i] = pc[i>>1] + byte(i&1)
    }
}

func PopCount1(x uint64) int {
    return int(
        pc[byte(x>>0)] +
        pc[byte(x>>8)] +
        pc[byte(x>>16)] +
        pc[byte(x>>24)] +
        pc[byte(x>>32)] +
        pc[byte(x>>40)] +
        pc[byte(x>>48)] +
        pc[byte(x>>56)])
}

func PopCount2(x uint64) int {
    var c byte
    for i := 0; i < 8; i++ {
        c += pc[byte(x>>(i*8))]
    }
    return int(c)
}

func PopCount3(x uint64) int {
    var c int
    for i := 0; i < 64; i++ {
        if x & 1 != 0 {
            c++
        }
        x >>= 1
    }
    return c
}

func PopCount4(x uint64) int {
    var c int
    for x > 0 {
        x &= x - 1
        c++
    }
    return c
}
