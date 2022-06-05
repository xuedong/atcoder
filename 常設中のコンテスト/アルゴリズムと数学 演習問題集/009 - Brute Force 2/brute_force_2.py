#!/usr/bin/env python3


def dp(nums, s):
    if len(nums) == 0:
        return False

    if nums[0] == s:
        return True
    elif nums[0] > s:
        return dp(nums[1:], s)
    else:
        return dp(nums[1:], s) or dp(nums[1:], s-nums[0])


if __name__ == "__main__":
    n, s = list(map(int, input().split()))
    nums = list(map(int, input().split()))

    flag = dp(nums, s)
    if flag:
        print("Yes")
    else:
        print("No")
