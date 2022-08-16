#!/usr/bin/env python3


if __name__ == "__main__":
    s1 = input()
    s2 = "atcoder"

    s1 = list(s1)
    s2 = list(s2)

    i, j = 0, 0
    ans = 0
    while i < 7:
        j = i
        while s1[j] != s2[i]:
            j += 1

        while i < j:
            s1[j], s1[j-1] = s1[j-1], s1[j]
            j -= 1
            ans += 1

        i += 1

    print(ans)
