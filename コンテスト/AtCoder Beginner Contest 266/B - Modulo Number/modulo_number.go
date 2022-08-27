package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int64
	fmt.Fscan(in, &n)

	ans := n % 998244353
	if ans >= 0 {
		fmt.Println(ans)
	} else {
		fmt.Println(ans + 998244353)
	}
}
