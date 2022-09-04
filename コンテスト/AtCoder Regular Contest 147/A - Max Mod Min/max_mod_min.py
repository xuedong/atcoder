#!/usr/bin/env python3

from collections import deque


if __name__ == "__main__":
    n = int(input())
    a = list(map(int, input().split()))

    a.sort()
    a = deque(a)
    ans = 0
    while len(a) > 1:
        curr = a.pop()
        if curr % a[0] > 0:
            a.appendleft(curr % a[0])
        ans += 1

    print(ans)
