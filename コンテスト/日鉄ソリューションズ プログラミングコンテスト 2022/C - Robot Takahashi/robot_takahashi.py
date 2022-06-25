#!/usr/bin/env python3

if __name__ == "__main__":
    n = int(input())
    s = input()
    weights = list(map(int, input().split()))

    pairs = [(weights[i], s[i]) for i in range(n)]
    pairs.sort()

    best = s.count("1")
    curr = best
    for i in range(n):
        if pairs[i][1] == "1":
            curr -= 1
        else:
            curr += 1
        if i < n-1:
            if pairs[i][0] != pairs[i+1][0]:
                best = max(curr, best)
        else:
            best = max(curr, best)
    best = max(best, s.count("0"))

    print(best)
