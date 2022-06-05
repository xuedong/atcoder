package main

import (
    "fmt"
)

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }

    return x
}

func main() {
    var n, x, y int
    fmt.Scanf("%d %d %d", &n, &x, &y)

  	mx, my, mxy := n / x, n / y, n / (x * y / gcd(x, y))

    fmt.Printf("%d\n", mx+my-mxy)
}
