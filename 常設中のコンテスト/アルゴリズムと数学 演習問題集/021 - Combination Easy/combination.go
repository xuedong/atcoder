package main

import (
    "fmt"
)

func fact(n int64) int64 {
  	var ans int64
    ans = 1
    var i int64
    for i = 1; i <= n; i++ {
        ans *= i
    }

    return ans
}

func main()  {
    var n, r int64
    fmt.Scanf("%v %d", &n, &r)

    ans := fact(n) / (fact(r) * fact(n-r))
    fmt.Printf("%v\n", ans)
}
