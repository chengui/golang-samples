package main

import (
    "fmt"
    "tree-sort/treesort"
)

func main() {
    arr := []int{1, 3, 2, 4, 0}
    treesort.Sort(arr)
    fmt.Println(arr)
}
