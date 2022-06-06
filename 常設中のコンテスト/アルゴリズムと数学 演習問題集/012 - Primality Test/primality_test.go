package main

import (
    "fmt"
    "math"
)

func isPrime(x int64) bool {
    if x <= 1 {
        return false
    }

    if x == 2 {
        return true
    }

	var i int64
  	for i = 2; i <= int64(math.Sqrt(float64(x))); i++ {
        if x % i == 0 {
            return false
        }
    }
    return true
}

func main()  {
    var n int64
    fmt.Scanf("%d", &n)

    if isPrime(n) {
        fmt.Printf("Yes")
    } else {
        fmt.Printf("No")
    }
}
