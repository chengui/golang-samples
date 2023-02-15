package main

import (
    "fmt"
    "pop-count/popcount"
)

func main() {
    fmt.Println(popcount.PopCount1(123456))
    fmt.Println(popcount.PopCount2(123456))
    fmt.Println(popcount.PopCount3(123456))
    fmt.Println(popcount.PopCount4(123456))
}
