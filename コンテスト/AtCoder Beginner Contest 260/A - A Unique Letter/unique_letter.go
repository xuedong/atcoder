package main

import (
    "bufio"
    "fmt"
    "os"
)

func main()  {
    in := bufio.NewReader(os.Stdin)

    var s string
    fmt.Fscan(in, &s)

    if s[0] == s[1] {
        if s[2] != s[1] {
            fmt.Println(string(s[2]))
        } else {
            fmt.Println(-1)
        }
    } else if s[0] == s[2] {
        fmt.Println(string(s[1]))
    } else {
        fmt.Println(string(s[0]))
    }
}
