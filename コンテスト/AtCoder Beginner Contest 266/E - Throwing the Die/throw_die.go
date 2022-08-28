package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func expectation(n int) float64 {
	if n == 1 {
		return 3.5
	}

	prev := expectation(n - 1)
	ans := 0.0
	for i := 1; i <= 6; i++ {
		ans += max(prev, float64(i)) / 6
	}
	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	fmt.Println(expectation(n))
}
