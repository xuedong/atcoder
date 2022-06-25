package main

import (
    "bufio"
    "fmt"
    "os"
)

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func main()  {
    in := bufio.NewReader(os.Stdin)

    var h1, h2, h3, w1, w2, w3 int
    fmt.Fscan(in, &h1, &h2, &h3, &w1, &w2, &w3)

    ans := 0
    for a11 := 1; a11 <= min(h1, w1)-2; a11++ {
        for a12 := 1; a12 <= h1-a11-1; a12++ {
            for a21 := 1; a21 <= w1-a11-1; a21++ {
                for a22 := 1; a22 <= min(h2-a21, w2-a12)-1; a22++ {
                    a31 := w1 - a11 - a21
                    a32 := w2 - a12 - a22
                    a13 := h1 - a11 - a12
                    a23 := h2 - a21 - a22
                    if a31 + a32 <= h3-1 && a13 + a23 <= w3-1 && h3-a31-a32 == w3-a13-a23 {
                        ans++
                    }
                }
            }
        }
    }

    fmt.Println(ans)
}
