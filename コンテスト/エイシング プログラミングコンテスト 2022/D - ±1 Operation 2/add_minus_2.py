#!/usr/bin/env python3

if __name__ == "__main__":
    n, q = list(map(int, input().split()))

    nums = list(map(int, input().split()))
    nums.sort()

    prefixes = [0] * (n+1)
    for i in range(1, n+1):
        prefixes[i] = prefixes[i-1] + nums[i-1]

    for j in range(q):
        target = int(input())

        left, right = 0, len(nums)-1
        while left <= right:
            mid = left + (right - left) // 2

            if nums[mid] < target:
                left = mid + 1
            else:
                right = mid - 1

        ans = left * target - prefixes[left] + prefixes[n] - prefixes[left] - (n - left) * target
        print(ans)
