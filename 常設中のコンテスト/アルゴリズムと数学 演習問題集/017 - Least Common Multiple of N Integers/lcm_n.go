package main

import (
    "bufio"
    "fmt"
    "os"
)

func gcd(a, b int) int  {
    if b == 0 {
        return a
    }
    return gcd(b, a % b)
}

func lcm(a, b int) int {
    return (b / gcd(a, b)) * a
}

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n int
    fmt.Fscan(in, &n)

    nums := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &nums[i])
    }

    ans := lcm(nums[0], nums[1])
    for i := 2; i < n; i++ {
        ans = lcm(ans, nums[i])
    }

    fmt.Printf("%v\n", ans)
}
