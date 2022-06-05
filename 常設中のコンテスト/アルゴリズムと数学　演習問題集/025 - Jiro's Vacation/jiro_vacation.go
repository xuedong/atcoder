package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)

    expectation := 0.0
    for i := 1; i <= n; i++ {
        var a float64
        fmt.Scanf("%f", &a)
        expectation += a / 3.0
    }
    for i := 1; i <= n; i++ {
        var b float64
        fmt.Scanf("%f", &b)
        expectation += b * 2.0 / 3.0
    }

    fmt.Printf("%.3f", expectation)
}
