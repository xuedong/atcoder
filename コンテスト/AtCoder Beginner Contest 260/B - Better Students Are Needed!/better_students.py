#!/usr/bin/env python3


def argsort(sequence):
    return sorted(range(len(sequence)), key=sequence.__getitem__)


if __name__ == "__main__":
    n, x, y, z = list(map(int, input().split()))
    a = list(map(int, input().split()))
    b = list(map(int, input().split()))

    c = [a[i]+b[i] for i in range(n)]

    used = set()

    indices_a = argsort([-a[i] for i in range(n)])
    for i in range(x):
        curr = indices_a[i]
        used.add(curr)

    indices_b = argsort([-b[i] for i in range(n)])
    j, count = 0, 0
    while count < y:
        curr = indices_b[j]
        if curr in used:
            j += 1
        else:
            used.add(curr)
            count += 1

    indices_c = argsort([-c[i] for i in range(n)])
    k, count = 0, 0
    while count < z:
        curr = indices_c[k]
        if curr in used:
            k += 1
        else:
            used.add(curr)
            count += 1

    for value in sorted(list(used)):
        print(value+1)
