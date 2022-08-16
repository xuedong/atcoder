package main

import (
    "bufio"
    "fmt"
    "os"
)

func isSubsequence(a []int, b []int) bool {
    i, j := 0, 0
    for i < len(a) && j < len(b) {
        if b[j] == a[i] {
            j++
        }
        i++
    }
    return j == len(b)
}

func main()  {
    in := bufio.NewReader(os.Stdin)

    var h1, w1 int
    fmt.Fscan(in, &h1, &w1)

    m1 := make([][]int, h1)
    for i := range m1 {
        m1[i] = make([]int, w1)
        for j := range m1[i] {
            fmt.Fscan(in, &m1[i][j])
        }
    }

    var h2, w2 int
    fmt.Fscan(in, &h2, &w2)

    m2 := make([][]int, h2)
    for i := range m2 {
        m2[i] = make([]int, w2)
        for j := range m2[i] {
            fmt.Fscan(in, &m2[i][j])
        }
    }

    i, j := 0, 0
    for i < len(m1[0]) && j < len(m2[0]) {
        if isSubsequence(m1[i], m2[j]) {
            j++
        }
        i++
    }
    if j == len(m2[0]) {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}
