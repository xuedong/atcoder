#!/usr/bin/env python3


if __name__ == "__main__":
    n, q = map(int, input().split())
    nums = [i for i in range(n+1)]
    inds = [i for i in range(n+1)]

    for i in range(q):
        num = int(input())
        
        idx1 = inds[num]
        idx2 = idx-1 if idx1 == n else idx+1
        num1, num2 = nums[idx1], nums[idx2]
          	
        nums[idx1], nums[idx2] = nums[idx2], nums[idx1]
        inds[num1], inds[num2] = inds[num2], inds[num1]

	print(' '.join(map(str, nums)))
