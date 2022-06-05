package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)

    blue, red := make([]float64, n), make([]float64, n)
    expBlue, expRed := 0.0, 0.0
    for i := 0; i < n; i++ {
        fmt.Scanf("%f", &blue[i])
        expBlue += blue[i]
    }
    for i := 0; i < n; i++ {
        fmt.Scanf("%f", &red[i])
        expRed += red[i]
    }

  	expBlue /= float64(n)
  	expRed /= float64(n)
    fmt.Printf("%.6f", expBlue+expRed)
}
