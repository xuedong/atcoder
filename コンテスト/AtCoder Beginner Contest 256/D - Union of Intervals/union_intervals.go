package main

import (
    "bufio"
    "fmt"
    "os"
	"sort"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	results := [][]int{intervals[0]}
	last := 0
	for i := 1; i < len(intervals); i++ {
		if results[last][1] >= intervals[i][0] {
			results[last][1] = max(results[last][1], intervals[i][1])
		} else {
			results = append(results, intervals[i])
			last++
		}
	}
	return results
}

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n int
    fmt.Fscan(in, &n)

	intervals := make([][]int, n)
	for i := 0; i < n; i++ {
		var left, right int
		fmt.Fscan(in, &left, &right)
		intervals[i] = make([]int, 2)
		intervals[i][0], intervals[i][1] = left, right
	}

	results := merge(intervals)

	for i := 0; i < len(results); i++ {
		fmt.Printf("%d %d\n", results[i][0], results[i][1])
	}
}
