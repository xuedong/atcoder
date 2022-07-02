package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func modulo(a, n int) int {
    if a % n < 0 {
        return a % n + n
    }
    return a % n
}

func helper(grid [][]int, i, j int, d1, d2 int) int {
    n := len(grid)

    ans := grid[i][j]
    for k := 1; k <= n-1; k++ {
        ans = 10 * ans + grid[modulo(i+k*d1, n)][modulo(j+k*d2, n)]
    }

    return ans
}

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n int
    fmt.Fscan(in, &n)

    grid := make([][]int, n)
    for row := range grid {
        grid[row] = make([]int, n)
    }

    for i := 0; i < n; i ++ {
        var curr string
        fmt.Fscan(in, &curr)
        for j := 0; j < n; j++ {
            grid[i][j], _ = strconv.Atoi(curr[j:j+1])
        }
    }

    neighbors := [8][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

    ans := 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            for _, neighbor := range neighbors {
                curr := helper(grid, i, j, neighbor[0], neighbor[1])
                if curr > ans {
                    ans = curr
                }
            }
        }
    }

    fmt.Println(ans)
}
