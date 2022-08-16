package main

import (
    "bufio"
    "fmt"
    "os"
)

func main()  {
    in := bufio.NewReader(os.Stdin)

    var l, r int
    fmt.Fscan(in, &l, &r)

    fmt.Println("atcoder"[l-1:r])
}
