package main

import (
    "fmt"
    "bytes"
    "strings"
)

func main() {
    fmt.Println(Comma3("123"))
    fmt.Println(Comma3("123456"))
    fmt.Println(Comma3("1234567"))
    fmt.Println(Comma3("-12"))
    fmt.Println(Comma3("-123"))
    fmt.Println(Comma3("-123456"))
    fmt.Println(Comma3("-1234567"))
    fmt.Println(Comma4("-12345.12345"))
}

func Comma1(s string) string {
    n := len(s)
    if n <= 3 || n == 4 && (s[0] == '+' || s[0] == '-') {
        return s
    }
    if s[0] == '+' || s[0] == '-' {
        return string(s[0]) + Comma1(s[1:])
    }
    return Comma1(s[:n-3]) + "," + s[n-3:]
}

func Comma2(s string) string {
    n := len(s)
    if n <= 3 || n == 4 && (s[0] == '+' || s[0] == '-') {
        return s
    }
    var r string
    if s[0] == '+' || s[0] == '-' {
        if (n - 1) % 3 == 0 {
            r = s[:4]
        } else {
            r = s[:1+(n-1)%3]
        }
    } else {
        if n % 3 == 0 {
            r = s[:3]
        } else {
            r = s[:n%3]
        }
    }
    for i := len(r); i < n; i += 3 {
        r += "," + s[i:i+3]
    }
    return r
}

func Comma3(s string) string {
    var buf bytes.Buffer
    i, n := 0, len(s)
    if s[0] == '+' || s[0] == '-' {
        buf.WriteByte(s[0])
        i++
    }
    if (n - i) % 3 == 0 {
        buf.WriteString(s[i:i+3])
        i += 3
    } else {
        buf.WriteString(s[i:i+(n-i)%3])
        i += (n-i)%3
    }
    for ; i < n; i += 3 {
        buf.WriteByte(',')
        buf.WriteString(s[i:i+3])
    }
    return buf.String()
}

func Comma4(s string) string {
    dot := strings.Index(s, ".")
    if dot >= 0 {
        return Comma1(s[:dot]) + s[dot:]
    }
    return Comma1(s)
}
