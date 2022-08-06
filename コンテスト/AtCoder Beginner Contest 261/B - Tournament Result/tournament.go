package main

import (
    "bufio"
    "fmt"
    "os"
)

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n int
    fmt.Fscan(in, &n)

    table := make([][]byte, n)
    for i := range table {
        table[i] = make([]byte, n)
    }

    for i := 0; i < n; i++ {
        var row string
        fmt.Fscan(in, &row)
        for j := 0; j < n; j++ {
            table[i][j] = row[j]
        }
    }

    flag := true
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            if table[i][j] == 'D' {
                if table[j][i] != 'D' {
                    flag = false
                    break
                }
            } else if table[i][j] == 'W' {
                if table[j][i] != 'L' {
                    flag = false
                    break
                }
            } else {
                if table[j][i] != 'W' {
                    flag = false
                    break
                }
            }
        }
    }

    if flag {
        fmt.Println("correct")
    } else {
        fmt.Println("incorrect")
    }
}
