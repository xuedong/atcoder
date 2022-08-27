package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var ax, ay int
	fmt.Fscan(in, &ax, &ay)
	var bx, by int
	fmt.Fscan(in, &bx, &by)
	var cx, cy int
	fmt.Fscan(in, &cx, &cy)
	var dx, dy int
	fmt.Fscan(in, &dx, &dy)

	a1 := cy - ay
	b1 := ax - cx
	c1 := a1*ax + b1*ay

	a2 := dy - by
	b2 := bx - dx
	c2 := a2*bx + b2*by

	determinant := float64(a1*b2 - a2*b1)

	if determinant == 0 {
		fmt.Println("No")
		return
	}

	x := float64(c1*b2-c2*b1) / determinant
	// y := (c2*a1 - c1*a2) / determinant

	if x > float64(min(ax, cx)) && x < float64(max(ax, cx)) && x > float64(min(bx, dx)) && x < float64(max(bx, dx)) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
