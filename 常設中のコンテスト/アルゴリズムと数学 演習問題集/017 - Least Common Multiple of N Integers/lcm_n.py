#!/usr/bin/env python3


def gcd(a, b):
    if b == 0:
        return a

    return gcd(b, a % b)

def lcm(a, b):
    return a * b // gcd(a, b)


if __name__ == "__main__":
    n = int(input())

    nums = list(map(int, input().split()))

    ans = lcm(nums[0], nums[1])
    for i in range(2, len(nums)):
        ans = lcm(ans, nums[i])

    print(ans)
