package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
)

func remove(slice []int, i int) []int {
    return append(slice[:i], slice[i+1:]...)
}

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n, x, y, z int
    fmt.Fscan(in, &n, &x, &y, &z)

    a := make([]int, n)
    b := make([]int, n)
    for i := range a {
        fmt.Fscan(in, &a[i])
    }
    for i := range b {
        fmt.Fscan(in, &b[i])
    }

    c := make([]int, n)
    for i := range c {
        c[i] = a[i] + b[i]
    }

    results := make([]int, 0)

    mapA := make(map[int][]int)
    mapB := make(map[int][]int)
    mapC := make(map[int][]int)
    for i := range a {
        mapA[a[i]] = append(mapA[a[i]], i)
        mapB[b[i]] = append(mapB[b[i]], i)
        mapC[c[i]] = append(mapC[c[i]], i)
    }

    sort.Sort(sort.Reverse(sort.IntSlice(a)))
    for k := 0; k < x; k++ {
        results = append(results, mapA[a[k]][0]+1)
        remove(b, mapA[a[k]][0])
        remove(c, mapA[a[k]][0])
        mapA[a[k]] = mapA[a[k]][1:]
    }

    sort.Sort(sort.Reverse(sort.IntSlice(b)))
    for k := 0; k < y; k++ {
        results = append(results, mapB[b[k]][0]+1)
        remove(c, mapB[b[k]][0])
        mapB[b[k]] = mapB[b[k]][1:]
    }

    sort.Sort(sort.Reverse(sort.IntSlice(c)))
    for k := 0; k < z; k++ {
        results = append(results, mapC[c[k]][0]+1)
        mapC[c[k]] = mapC[c[k]][1:]
    }

    sort.Ints(results)
    for i := range results{
        fmt.Println(results[i])
    }
}
