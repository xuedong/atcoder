package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var h, w int
  	fmt.Scanf("%d %d", &h, &w)

  	i1, j1, i2, j2 := -1, -1, -1, -1
  	for i := 1; i <= h; i++ {
      	var row string
      	fmt.Scanf("%v", &row)

      	for j := 0; j < w; j++ {
          	if row[j] == 'o' {
              	if i1 == -1 {
                	i1, j1 = i, j
                } else {
                  	i2, j2 = i, j
                }
            }
        }
    }

  	ans := abs(j2-j1) + abs(i2-i1)
  	fmt.Printf("%d\n", ans)
}
