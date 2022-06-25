package main

import (
    "bufio"
    "fmt"
    "os"
    "math"
)

func main()  {
    in := bufio.NewReader(os.Stdin)

    var ax, ay, bx, by, cx, cy float64
    fmt.Fscan(in, &ax, &ay, &bx, &by, &cx, &cy)

    var res float64
	if (bx - ax) * (bx - cx) + (by - ay) * (by - cy) < 0 {
		res = math.Sqrt((bx - ax) * (bx - ax) + (by - ay) * (by - ay))
	} else if (cx - ax) * (cx - bx) + (cy - ay) * (cy - by) < 0 {
		res = math.Sqrt((cx - ax) * (cx - ax) + (cy - ay) * (cy - ay))
	} else {
		res = math.Abs((cx - bx) * (ay - by) - (cy - by) * (ax - bx)) / math.Sqrt((cx - bx) * (cx - bx) + (cy - by) * (cy - by))
	}

	fmt.Println(res)
}
