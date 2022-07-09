package main

import (
    "bufio"
    "fmt"
    "os"
    "math"
)

func union(i, j int, parents []int) {
	u1, u2 := find(i, parents), find(j, parents)
	if u1 != u2 {
		parents[u2] = u1
	}
}

func find(i int, parents []int) int {
	if parents[i] == i {
		return i
	}
	return find(parents[i], parents)
}

func isConnected(x1, y1, r1 float64, x2, y2, r2 float64) bool {
    distance := math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
    return distance <= r1 + r2 && distance >= math.Abs(r1 - r2)
}

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n int
    fmt.Fscan(in, &n)

    var sx, sy, tx, ty float64
    fmt.Fscan(in, &sx, &sy, &tx, &ty)

    circles := make([][3]float64, n)
    for i := range circles {
        fmt.Fscan(in, &circles[i][0], &circles[i][1], &circles[i][2])
    }

	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}

    start, end := -1, -1
    for i := 0; i < n; i++ {
        xi, yi, ri := circles[i][0], circles[i][1], circles[i][2]
        if (sx - xi) * (sx - xi) + (sy - yi) * (sy - yi) == ri * ri {
            start = i
        }
        if (tx - xi) * (tx - xi) + (ty - yi) * (ty - yi) == ri * ri {
            end = i
        }
        for j := i+1; j < n; j++ {
			xj, yj, rj := circles[j][0], circles[j][1], circles[j][2]
			if isConnected(xi, yi, ri, xj, yj, rj) {
				union(i, j, parents)
			}
        }
    }

	if find(start, parents) == find(end, parents) {
		fmt.Printf("Yes")
	} else {
		fmt.Printf("No")
	}
}
