#!/usr/bin/env python3


if __name__ == "__main__":
    n, q = map(int, input().split())
    nums = [i for i in range(1, n+1)]

    for i in range(q):
        num = int(input())
        idx = nums.index(num)
        if idx == n-1:
            nums[idx], nums[idx-1] = nums[idx-1], nums[idx]
        else:
            nums[idx], nums[idx+1] = nums[idx+1], nums[idx]

	print(' '.join(map(str, nums)))

