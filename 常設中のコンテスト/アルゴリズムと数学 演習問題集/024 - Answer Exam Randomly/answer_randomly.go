package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)

    expectation := 0.0
    for i := 1; i <= n; i++ {
        var p, q float64
        fmt.Scanf("%f %f", &p, &q)
        expectation += q / p
    }

    fmt.Printf("%.6f", expectation)
}
