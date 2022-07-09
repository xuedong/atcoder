#include<bits/stdc++.h>
#define rep(i, a, b) for (int i = a; i < b; i++)
#define rrep(i, a, b) for (int i = a; i >= b; i--)
#define fore(i, a) for (auto &i: a)
#define all(x) (x).begin(), (x).end()
//#pragma GCC optimize ("-O3")
using namespace std; void _main(); int main() { cin.tie(0); ios::sync_with_stdio(false); _main(); }
typedef long long ll; const int inf = INT_MAX / 2; const ll infl = 1LL << 60;
template<class T>bool chmax(T& a, const T& b) { if (a < b) { a = b; return 1; } return 0; }
template<class T>bool chmin(T& a, const T& b) { if (b < a) { a = b; return 1; } return 0; }

int N;
ll sx, sy, tx, ty;
ll x[3010], y[3010], r[3010];

int getCircleOn(int px, int py) {
    rep(i, 0, N) {
        ll dist2 = abs(px - x[i]) * abs(px - x[i]) + abs(py - y[i]) * abs(py - y[i]);
        if (dist2 == r[i] * r[i]) return i;
    }
    return -1;
}

#define YES "Yes"
#define NO "No"
string solve() {
    int start = getCircleOn(sx, sy);
    int end = getCircleOn(tx, ty);

    //
    // makeEdges
    // ================================
    vector<int> E[3010];
    rep(i, 0, N) rep(j, i + 1, N) {
        ll dist2 = (x[i] - x[j]) * (x[i] - x[j]) + (y[i] - y[j]) * (y[i] - y[j]);

        if ((r[i] + r[j]) * (r[i] + r[j]) < dist2) continue;
        if (dist2 < abs(r[i] - r[j]) * abs(r[i] - r[j])) continue;

        E[i].push_back(j);
        E[j].push_back(i);
    }

    //
    // checkReachable
    // ================================
    queue<int> que;
    set<int> visited;

    que.push(start);
    visited.insert(start);

    while (!que.empty()) {
        int cu = que.front();
        que.pop();

        if (cu == end) return YES;

        fore(to, E[cu]) if (!visited.count(to)) {
            que.push(to);
            visited.insert(to);
        }
    }

    return NO;
}

void _main() {
    cin >> N >> sx >> sy >> tx >> ty;
    rep(i, 0, N) cin >> x[i] >> y[i] >> r[i];
    cout << solve() << endl;
}
