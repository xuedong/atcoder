package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	pits := make([]int, 100001)
	sizes := make([]int64, 100001)
	for i := 0; i < n; i++ {
		var time, pit int
		var size int64
		fmt.Fscan(in, &time, &pit, &size)
		pits[time] = pit
		sizes[time] = size
	}

	dp := make([][]int64, 100001)
	for row := range dp {
		dp[row] = make([]int64, 5)
		for col := range dp[row] {
			dp[row][col] = -1000000000000
		}
	}
	dp[0][0] = 0

	for t := 1; t < 100001; t++ {
		for i := 0; i < 5; i++ {
			dp[t][i] = dp[t-1][i]
			if i != 0 {
				dp[t][i] = max(dp[t-1][i-1], dp[t][i])
			}
			if i != 4 {
				dp[t][i] = max(dp[t-1][i+1], dp[t][i])
			}
		}
		dp[t][pits[t]] += sizes[t]
	}

	var ans int64
	ans = 0
	for i := 0; i < 5; i++ {
		ans = max(ans, dp[100000][i])
	}
	fmt.Printf("%d\n", ans)
}
