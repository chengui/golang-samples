package xkcd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

type XkcdInfo struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

type Xkcd struct {
	cache map[string]string
}

func NewXkcd() *Xkcd {
	return &Xkcd{
		cache: make(map[string]string),
	}
}

func (x *Xkcd) fetch(start, end int) error {
	result := make(chan string, 10)
	go func(result <-chan string) {
		cache := fmt.Sprintf("xkcd-%d-%d.csv", start, end)
		f, err := os.Create(cache)
		if err != nil {
			return
		}
		defer f.Close()
		for rc := range result {
			f.Write([]byte(rc))
		}
	}(result)
	wg := &sync.WaitGroup{}
	for i := start; i < end; i++ {
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		wg.Add(1)
		go func(url string, result chan<- string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				return
			}
			var info XkcdInfo
			if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
				return
			}
			result <- fmt.Sprintf("%s,%s\n", strings.ToLower(info.Title), info.Img)
		}(url, result)
	}
	wg.Wait()
	close(result)
	return nil
}

func (x *Xkcd) Load(start, end int) error {
	cache := fmt.Sprintf("xkcd-%d-%d.csv", start, end)
	if _, err := os.Stat(cache); os.IsNotExist(err) {
		x.fetch(start, end)
	}
	file, err := os.Open(cache)
	if err != nil {
		return err
	}
	defer file.Close()
	input := bufio.NewScanner(file)
	for input.Scan() {
		line := input.Text()
		kvs := strings.Split(line, ",")
		x.cache[kvs[0]] = kvs[1]
	}
	return nil
}

func (x *Xkcd) Find(key string) string {
	return x.cache[key]
}
