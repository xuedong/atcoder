package main

import (
    "bufio"
    "fmt"
    "os"
)

func aux(n, k, start int, curr []int, res *[][]int) {
    if len(curr) == k {
        temp := make([]int, k)
        copy(temp, curr)
        *res = append(*res, temp)
        return
    }

    for i := start; i <= n; i++ {
        curr = append(curr, i)
        aux(n, k, i+1, curr, res)
        curr = curr[:len(curr)-1]
    }
}

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n, m int
    fmt.Fscan(in, &n, &m)

    curr := make([]int, 0)
    res := [][]int{}

    aux(m, n, 1, curr, &res)

    for i := range res {
        for j := range res[i] {
            fmt.Printf("%d ", res[i][j])
        }
        fmt.Println()
    }
}
