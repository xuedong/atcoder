#!/usr/bin/env python3

import heapq as hq


if __name__ == "__main__":
	n, l = map(int, input().split())
	nums = map(int, input().split())

	queue = []
	hq.heapify(queue)

	rest = l
	for num in nums:
		hq.heappush(queue, num)
		rest -= num
	if rest > 0:
		hq.heappush(queue, rest)

	ans = 0
	while len(queue) > 1:
		x = hq.heappop(queue)
		y = hq.heappop(queue)
		hq.heappush(queue, x+y)
		ans += x + y

	print(ans)
