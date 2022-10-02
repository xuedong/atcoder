#include <iostream>
#include <vector>
#include <queue>

using namespace std;

int N;
vector<int> graph[100000];
int dist[100000];

void getDist(int start) {
    for (int i = 0; i < N; i++) {
        dist[i] = -1;
    }

    queue<int> q;
    q.push(start);
    dist[start] = 0;
    while (!q.empty()) {
        int next = q.front();
        q.pop();
        for (int i = 0; i < graph[next].size(); i++) {
            int to = graph[next][i];
            if (dist[to] != -1) {
                continue;
            }
            dist[to] = dist[next] + 1;
            q.push(to);
        }
    }
}

int main() {
    cin >> N;

    for (int i = 0; i < N-1; i++) {
        int a, b;
        cin >> a >> b;
        a--;
        b--;
        graph[a].push_back(b);
        graph[b].push_back(a);
    }

    getDist(0);
    int maxDist = 0;
    int maxDistNode = 0;
    for (int i = 0; i < N; i++) {
        if (dist[i] > maxDist) {
            maxDist = dist[i];
            maxDistNode = i;
        }
    }

    getDist(maxDistNode);
    maxDist = 0;
    for (int i = 0; i < N; i++) {
        maxDist = max(maxDist, dist[i]);
    }

    cout << maxDist + 1 << endl;
}
