package main

import (
	"fmt"
)

func gcd(a, b int64) int64 {
  	for b != 0 {
      	t := b
        b = a % b
        a = t
    }
    return a
}

func lcm(a, b int64) int64 {
    result := a * b / gcd(a, b)
    return result
}

func main() {
	var n, a, b int64
  	fmt.Scanf("%d %d %d", &n, &a, &b)
  	c := lcm(a, b)

  	sum := n * (n+1) / 2
  	sumA := (a + n / a * a) * (n / a) / 2
  	sumB := (b + n / b * b) * (n / b) / 2
  	sumC := (c + n / c * c) * (n / c) / 2

  	result := sum - (sumA + sumB) + sumC
  	fmt.Printf("%d\n", result)
}
