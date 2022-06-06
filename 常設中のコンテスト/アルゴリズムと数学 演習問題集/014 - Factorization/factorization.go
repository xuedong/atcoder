package main

import (
    "fmt"
    "math"
)

func main()  {
    var n int64
    fmt.Scanf("%d", &n)

    var i int64
  	for i = 2; i <= int64(math.Sqrt(float64(n))); i++ {
        for n % i == 0 && n != 0 {
            n /= i
          	fmt.Printf("%d ", i)
        }

        if n == 0 {
            break
        }
    }

    if n != 1 {
        fmt.Printf("%d", n)
    }
}
