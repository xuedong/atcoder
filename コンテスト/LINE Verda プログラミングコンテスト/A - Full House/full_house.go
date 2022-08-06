package main

import (
    "bufio"
    "fmt"
    "os"
)

func main()  {
    in := bufio.NewReader(os.Stdin)

    var a, b, c, d, e int
    fmt.Fscan(in, &a, &b, &c, &d, &e)

    if a == b {
        if c == a && d == e {
            fmt.Println("Yes")
        } else if d == a && c == e {
            fmt.Println("Yes")
        } else if e == a && c == d {
            fmt.Println("Yes")
        } else if c == d && d == e {
            fmt.Println("Yes")
        } else {
            fmt.Println("No")
        }
    } else {
        if c == a && d == a && e == b {
            fmt.Println("Yes")
        } else if c == a && e == a && d == b {
            fmt.Println("Yes")
        } else if d == a && e == a && c == b {
            fmt.Println("Yes")
        } else if c == b && d == b && e == a {
            fmt.Println("Yes")
        } else if c == b && e == b && d == a {
            fmt.Println("Yes")
        } else if d == b && e == b && c == a {
            fmt.Println("Yes")
        } else {
            fmt.Println("No")
        }
    }
}
