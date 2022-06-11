package main

import (
    "bufio"
    "fmt"
    "os"
)

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n, x, y int
    fmt.Fscan(in, &n, &x, &y)

    if x < 0 {
        x = -x
    }
    if y < 0 {
        y = -y
    }

  	if n - x - y >= 0 && (n - x - y) % 2 == 0 {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}
