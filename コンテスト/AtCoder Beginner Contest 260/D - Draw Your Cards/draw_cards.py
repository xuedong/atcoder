#!/usr/bin/env python3

from bisect import bisect_left


if __name__ == "__main__":
    n, k = map(int, input().split())
    p = list(map(int, input().split()))

    top_cards = []
    counts = []
    stacks = []
    results = [-1] * n

    for i in range(n):
        idx = bisect_left(top_cards, p[i])
        if idx == len(top_cards):
            top_cards.append(p[i])
            counts.append(1)
            stacks.append([p[i]])
        else:
            top_cards[idx] = p[i]
            counts[idx] += 1
            stacks[idx].append(p[i])

        if counts[idx] == k:
            top_cards.pop(idx)
            counts.pop(idx)
            for j in stacks[idx]:
                results[j-1] = i+1
            stacks.pop(idx)

    for i in range(n):
        print(results[i])
