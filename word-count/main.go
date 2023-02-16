package main

import (
    "os"
    "word-count/wordcount"
)

func main() {
    wordcount.CountWord(os.Stdin)
}
