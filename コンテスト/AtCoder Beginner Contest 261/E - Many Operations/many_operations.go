package main

import (
    "bufio"
    "fmt"
    "os"
)

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n, c int
    fmt.Fscan(in, &n, &c)

    ops := make([][2]int, n)
    for i := range ops {
        fmt.Fscan(in, &ops[i][0], &ops[i][1])
    }

    results := make([]int, n)
	for k := 0; k < 30; k++ {
		curr := (c >> k) & 1
		for i := 0; i < n; i++ {
			x := (ops[i][1] >> k) & 1
			
			if ops[i][0] == 1 {
				curr &= x
			}
			if ops[i][0] == 2 {
				curr |= x
			}
			if ops[i][0] == 3 {
				curr ^= x
			}

			results[i] |= (curr << k)
		}
	}

	for i := range results {
		fmt.Printf("%d\n", results[i])
	}
}
