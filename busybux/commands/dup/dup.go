package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func Main(w io.Writer, args []string) {
    dup2(w, args)
}

func dup1(w io.Writer, files []string) {
    countLines := func(file *os.File, counts map[string]int) {
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            counts[scanner.Text()]++
        }
    }
    counts := make(map[string]int)
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(w, "dup: %v\n", err)
                continue
            }
            countLines(f, counts)
            f.Close()
        }
    for line, count := range counts {
        if count > 1 {
            fmt.Fprintf(w, "%d\t%q\n", count, line)
        }
    }
}

func dup2(w io.Writer, files []string) {
    counts := make(map[string]int)
    if len(files) == 0 {
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
            counts[scanner.Text()]++
        }
    } else {
        for _, file := range files {
            data, err := ioutil.ReadFile(file)
            if err != nil {
                fmt.Fprintf(w, "dup: %v\n", err)
                continue
            }
            for _, line := range strings.Split(string(data), "\n") {
                counts[line]++
            }
        }
    }
    for line, count := range counts {
        if count > 1 {
            fmt.Fprintf(w, "%d\t%q\n", count, line)
        }
    }
}
