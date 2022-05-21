#!/usr/bin/env python3

import heapq as hq
from typing import List


def road_reduction(graph: List[List[int]], n: int) -> List[int]:
    visited = [False for _ in range(n)]

    queue = graph[0]
    visited[0] = True
    hq.heapify(queue)

    results = []
    while queue:
        distance, node, idx = hq.heappop(queue)

        if visited[node]:
            continue
        results.append(idx)
        visited[node] = True

        for neighbor in graph[node]:
            neighbor_distance, neighbor_node, neighbor_idx = neighbor
            if visited[neighbor_node]:
                continue
            hq.heappush(queue, (distance + neighbor_distance, neighbor_node, neighbor_idx))

    return results


if __name__ == "__main__":
    n, m = map(int, input().split())

    graph = [[] for _ in range(n)]
    for i in range(1, m+1):
        u, v, distance = map(int, input().split())
        graph[u-1].append((distance, v-1, i))
        graph[v-1].append((distance, u-1, i))

    results = road_reduction(graph, n)
    print(*results)
