package main

import (
    "fmt"
)

func main()  {
    var a, b int
    fmt.Scanf("%d %d", &a, &b)

    for b > 0 {
        a, b = b, a % b
    }

    fmt.Printf("%d\n", a)
}
