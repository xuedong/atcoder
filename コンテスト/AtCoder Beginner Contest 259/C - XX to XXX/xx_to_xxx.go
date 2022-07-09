package main

import (
    "bufio"
    "fmt"
    "os"
)

func main()  {
    in := bufio.NewReader(os.Stdin)

    var s, t string
    fmt.Fscan(in, &s, &t)

    i, j := 0, 0
    for i < len(s) && j < len(t) {
        countT := 0
        for j < len(t) && s[i] == t[j] {
            countT++
            j++
        }

        if countT == 0 {
            fmt.Printf("No")
            return
        }

        countS := 0
        for i < len(s) && s[i] == t[j-1] {
            countS++
            i++
        }

        if countS == 1 && countT >= 2 {
            fmt.Printf("No")
            return
        } else if countS > countT {
            fmt.Printf("No")
            return
        }
    }

    if i < len(s) || j < len(t) {
        fmt.Printf("No")
    } else {
        fmt.Printf("Yes")
    }
}
