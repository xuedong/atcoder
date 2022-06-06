package main

import (
    "fmt"
    "math"
)

func isPrime(x int) bool {
    if x <= 1 {
        return false
    }

    if x == 2 {
        return true
    }

  	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
        if x % i == 0 {
            return false
        }
    }
    return true
}

func main()  {
    var n int
    fmt.Scanf("%d", &n)

    for i := 2; i <= n; i++ {
        if isPrime(i) {
            fmt.Printf("%d ", i)
        }
    }
    fmt.Println()
}
