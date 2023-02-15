package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
    "time"
)

func main() {
    if false {
        for _, url := range os.Args[1:] {
            if err := fetch2(url); err != nil {
                os.Exit(1)
            }
        }
    } else {
        start := time.Now()
        ch := make(chan string)
        for _, url := range os.Args[1:] {
            go fetch3(url, ch)
        }
        for range os.Args[1:] {
            fmt.Println(<-ch)
        }
        fmt.Printf("%.2fs elapsed.\n", time.Since(start).Seconds())
    }
}

func fetch1(url string) error {
    if !strings.HasPrefix(url, "http://") {
        url = "http://" + url
    }
    resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        return err
    }
    data, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        return err
    }
    fmt.Printf("Status: %v\n", resp.Status)
    fmt.Printf("%s", data)
    return nil
}

func fetch2(url string) error {
    if !strings.HasPrefix(url, "http://") {
        url = "http://" + url
    }
    resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        return err
    }
    fmt.Printf("Status: %v\n", resp.Status)
    _, err = io.Copy(os.Stdout, resp.Body)
    resp.Body.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        return err
    }
    return nil
}

func fetch3(url string, ch chan<- string) {
    start := time.Now()
    if !strings.HasPrefix(url, "http://") {
        url = "http://" + url
    }
    resp, err := http.Get(url)
    if err != nil {
        ch<- fmt.Sprint(err)
        return
    }
    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()
    if err != nil {
        ch<- fmt.Sprint(err)
        return
    }
    ch<- fmt.Sprintf("%.2fs %7d %s", time.Since(start).Seconds(), nbytes, url)
}
