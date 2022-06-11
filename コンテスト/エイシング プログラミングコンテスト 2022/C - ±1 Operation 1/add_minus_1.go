package main

import (
	"fmt"
)

func binarySearch(target int64, a int64, d int64, n int64) int64 {
	var left, right int64
	left, right = 0, n - 1

	for left <= right {
		mid := left + (right - left) / 2

		if a + mid * d < target {
			if d > 0 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if d > 0 {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return right
}

func main()  {
	var x, a, d, n int64
	fmt.Scanf("%v %v %v %v", &x, &a, &d, &n)

	id := binarySearch(x, a, d, n)
  	if id == -1 {
      	if d > 0 {
          	fmt.Printf("%v", a - x)
        } else {
          	fmt.Printf("%v", x - a)
        }
      	return
    }
	if id == n-1 {
      	if d > 0 {
          	fmt.Printf("%v", x - (a + (n - 1) * d))
        } else {
          	fmt.Printf("%v", a + (n - 1) * d - x)
        }
      	return
    }

	var diff1, diff2 int64
	if d > 0 {
		diff1 = x - (a + id * d)
		diff2 = (a + (id + 1) * d) - x
	} else {
      	diff1 = (a + id * d) - x
      	diff2 = x - (a + (id + 1) * d)
	}

	var ans int64

	if diff1 < diff2 {
		ans = diff1
	} else {
		ans = diff2
	}

	fmt.Printf("%v\n", ans)
}
