package main

import (
    "bufio"
    "fmt"
    "os"
)

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n, k int
    fmt.Fscan(in, &n, &k)

    set := map[int]bool{}

    if k > n/2 {
        fmt.Println(-1)
    } else {
        var i int
        t := 1
        for i = 1; i <= n; i++ {
            _, ok := set[i+k]
            if i+k <= n && !ok && i+k > n-k {
                fmt.Printf("%d ", i+k)
            } else {
                for {
                    if _, used := set[t]; !used {
                        break
                    }
                    t++
                }
                if t > i-k {
					set[i+k] = true
					fmt.Printf("%d ", i+k)
				} else {
					set[t] = true
					fmt.Printf("%d ", t)
				}
			}
        }
    }
}
