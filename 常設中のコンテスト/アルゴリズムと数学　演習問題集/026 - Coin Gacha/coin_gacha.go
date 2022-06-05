package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)

    expectation := 0.0
    for i := 1; i <= n-1; i++ {
        expectation += float64(n) / float64(i)
    }

    fmt.Printf("%.6f", expectation+1)
}
