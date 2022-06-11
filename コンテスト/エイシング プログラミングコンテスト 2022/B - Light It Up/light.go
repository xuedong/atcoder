package main

import (
    "fmt"
    "math"
)

func main()  {
    var n, k int
    fmt.Scanf("%d %d", &n, &k)

    centers := make([]int, k)
    for i := 0; i < k; i++ {
        fmt.Scanf("%d", &centers[i])
    }

    points := make([][]float64, n)
    for i := 0; i < n; i++ {
        points[i] = make([]float64, 2)
        fmt.Scanf("%f %f", &points[i][0], &points[i][1])
    }

    distances := make([]float64, n)
    for i := 0; i < n; i++ {
        distance := math.Inf(1)
        for _, k := range centers {
            if i == k-1 {
                distance = -1.0
                break
            }

            curr := math.Sqrt(math.Pow(points[i][0] - points[k-1][0], 2) + math.Pow(points[i][1] - points[k-1][1], 2))
          	if curr < distance {
                distance = curr
            }
        }

        if distance == -1.0 {
            continue
        }
        distances[i] = distance
    }

    ans := 0.0
    for i := 0; i < n; i++ {
        if distances[i] > ans {
            ans = distances[i]
        }
    }
    fmt.Printf("%.5f", ans)
}
