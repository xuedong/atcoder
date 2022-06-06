package main

import (
    "fmt"
)

func main() {
    var n int64
    fmt.Scanf("%d", &n)

    var ans int64
    ans = 1
  	var i int64
    for i = 1; i <= n; i++ {
        ans *= i
    }

    fmt.Printf("%d\n", ans)
}
