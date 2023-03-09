package echo

import "os"
import "fmt"
import "strings"

func Main(w io.Writer, args []string) {
    echo1(w, args)
}

func echo1(w, args []string) {
    var s, sep string
    for i := 0; i < len(args); i++ {
        s += sep + args[i]
        sep = " "
    }
    fmt.Fprintln(w, s)
}

func echo2(args []string) {
    s, sep := "", ""
    for _, arg := range args {
        s += sep + arg
        sep = " "
    }
    fmt.Fprintln(w, s)
}
