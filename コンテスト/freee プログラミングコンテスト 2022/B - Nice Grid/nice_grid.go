package main

import (
    "bufio"
    "fmt"
    "os"
)

func main()  {
    in := bufio.NewReader(os.Stdin)

    var r, c int
    fmt.Fscan(in, &r, &c)

    if r == 1 || r == 15 {
        fmt.Println("black")
    }
    if r == 3 || r == 13 {
        if c != 2 && c != 14 {
            fmt.Println("black")
        } else {
            fmt.Println("white")
        }
    }
    if r == 5 || r == 11 {
        if c != 2 && c != 4 && c != 14 && c != 12 {
            fmt.Println("black")
        } else {
            fmt.Println("white")
        }
    }
    if r == 7 || r == 9 {
        if c != 2 && c != 4 && c != 6 && c != 14 && c != 12 && c != 10 {
            fmt.Println("black")
        } else {
            fmt.Println("white")
        }
    }
    if r == 8 {
        if c % 2 == 1 {
            fmt.Println("black")
        } else {
            fmt.Println("white")
        }
    }
    if r == 2 || r == 14 {
        if c == 1 || c == 15 {
            fmt.Println("black")
        } else {
            fmt.Println("white")
        }
    }
    if r == 4 || r == 12 {
        if c == 1 || c == 3 || c == 15 || c == 13 {
            fmt.Println("black")
        } else {
            fmt.Println("white")
        }
    }
    if r == 6 || r == 10 {
        if c == 1 || c == 3 || c == 5 || c == 15 || c == 13 || c == 11 {
            fmt.Println("black")
        } else {
            fmt.Println("white")
        }
    }
}
