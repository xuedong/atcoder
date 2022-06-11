#!/usr/bin/env python3


if __name__ == "__main__":
    n = int(input())

    nums = list(map(int, input().split()))

    counts = [0] * 100000
    for num in nums:
        counts[num] += 1

    ans = 0
    for i in range(1, 50000):
        ans += counts[i] * counts[100000 - i]
    ans += counts[50000] * (counts[50000] - 1) // 2

    print(ans)
