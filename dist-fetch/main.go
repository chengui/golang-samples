package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
)

func main() {
    for _, url := range os.Args[1:] {
        if !strings.HasPrefix(url, "http://") {
            url = "http://" + url
        }
        if err := fetch2(url); err != nil {
            os.Exit(1)
        }
    }
}

func fetch1(url string) error {
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
    resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        return err
    }
    fmt.Printf("Status: %v\n", resp.Status)
    _, err = io.Copy(os.Stdout, resp.Body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        return err
    }
    return nil
}
