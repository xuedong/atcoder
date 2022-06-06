package main

import (
    "fmt"
    "math"
)

func main()  {
    var n int64
    fmt.Scanf("%d", &n)

    divisors := make([]int64, 0)
    var i int64
  	for i = 1; i <= int64(math.Sqrt(float64(n))); i++ {
        if n % i == 0 {
            divisors = append(divisors, n/i)
          	fmt.Println(i)
        }
    }

  	for j := len(divisors)-1; j >= 0; j-- {
        fmt.Println(divisors[j])
    }
}
