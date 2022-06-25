from collections import deque


if __name__ == "__main__":
    n = int(input())
    xs = list(map(int, input().split()))
    cs = list(map(int, input().split()))

    visited = set()
    ans = 0
    for i in range(1, n+1):
        temp = i
        if temp in visited:
            continue

        queue = deque()
        while temp not in visited:
            queue.append(temp)
            visited.add(temp)
            temp = xs[temp-1]

        while queue:
            if queue[0] != temp:
                queue.popleft()
            else:
                break

        curr = []
        if len(queue) > 0:
            for i in queue:
                curr.append(cs[i-1])
            ans += min(curr)

    print(ans)
