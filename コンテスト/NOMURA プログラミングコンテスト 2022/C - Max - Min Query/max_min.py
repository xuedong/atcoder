if __name__ == "__main__":
    q = int(input())
    multiset = {}

    for _ in range(q):
        query = input().split()

        if query[0] == '1':
            x = int(query[1])
            if x in multiset:
                multiset[x] += 1
            else:
                multiset[x] = 0

        if query[0] == '2':
            x, c = int(query[1]), int(query[2])
            if x in multiset:
                multiset[x] -= c
                if multiset[x] <= 0:
                    multiset.pop(x)

        if query[0] == '3':
            print(max(multiset.keys()) - min(multiset.keys()))
