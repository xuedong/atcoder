package main

import (
    "fmt"
)

func main()  {
    var r, c int
    fmt.Scanf("%d %d", &r, &c)

    var a11, a12 int
    var a21, a22 int
    fmt.Scanf("%d %d", &a11, &a12)
    fmt.Scanf("%d %d", &a21, &a22)

    if r == 1 && c == 1 {
        fmt.Printf("%d", a11)
    } else if r == 1 && c == 2 {
        fmt.Printf("%d", a12)
    } else if r == 2 && c == 1 {
        fmt.Printf("%d", a21)
    } else {
        fmt.Printf("%d", a22)
    } 
}
