package main

import (
	"fmt"
)

func main() {
	var n int64
  	fmt.Scanf("%d", &n)

  	var ans int64
	ans = 0

	var i int64
	for i = 1; i <= n; i++ {
		var rest int64
		rest = i
		var factor int64
		for factor = 2; factor * factor <= rest; factor++ {
			for rest % (factor * factor) == 0 {
				rest /= factor * factor
			}
		}

		for factor = 1; factor * factor * rest <= n; factor++ {
			ans++
		}
	}

	fmt.Printf("%d", ans)
}
