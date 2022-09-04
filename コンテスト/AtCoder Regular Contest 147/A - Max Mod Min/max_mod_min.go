package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	sort.Ints(a)
	ans := 0
	for a[0] > 1 {
		curr := a[len(a)-1] % a[0]
		if curr == 0 {
			a = a[:len(a)-1]
		} else {
			a = a[:len(a)-1]
			a = append([]int{curr}, a...)
		}
		ans++
	}

	fmt.Println(ans + len(a) - 1)
}
