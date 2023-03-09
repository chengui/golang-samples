package wc

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func Main(w io.Writer, args []string) error {
	for _, arg := range args {
		f, _ := os.Open(arg)
		CountWord(w, f)
	}
	return nil
}

func CountRune(rd io.Reader) {
	countLetters, countNumbers := 0, 0
	countInvalid, countLines := 0, 0
	countChars := make(map[rune]int)
	countLens := [utf8.UTFMax + 1]int{}
	reader := bufio.NewReader(rd)
	for {
		r, n, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("error: %v", err)
			continue
		}
		if r == unicode.ReplacementChar && n == 1 {
			countInvalid++
			continue
		}
		if r == rune('\n') {
			countLines++
		}
		if unicode.IsLetter(r) {
			countLetters++
		}
		if unicode.IsNumber(r) {
			countNumbers++
		}
		countChars[r]++
		countLens[n]++
	}
	fmt.Printf("Lines: %d\n", countLines)
	fmt.Printf("Invalid: %d\n", countInvalid)
	fmt.Printf("Letters: %d\n", countLetters)
	fmt.Printf("Numbers: %d\n", countNumbers)
	fmt.Printf("len\tcnt\n")
	for i := 1; i < len(countLens); i++ {
		fmt.Printf("%d\t%d\n", i, countLens[i])
	}
	fmt.Printf("char\tcnt\n")
	for c, n := range countChars {
		fmt.Printf("%q\t%d\n", c, n)
	}
}

func CountWord(w io.Writer, rd io.Reader) {
	countWords := make(map[string]int)
	scanner := bufio.NewScanner(rd)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		countWords[word]++
	}
	fmt.Printf("word\tcnt\n")
	for c, n := range countWords {
		fmt.Fprintf(w, "%q\t%d\n", c, n)
	}
}
