#include <iostream>
using namespace std;

long long K, N, L;
long long A[1 << 18];

bool check(long long x) {
    long long count = 0;
    long long prev = 0;
    for (long long i = 0; i < N; i++) {
        if (A[i] - prev >= x && L - A[i] >= x) {
            count++;
            prev = A[i];
        }
    }
    return count >= K;
}

int main() {
    cin >> N >> L;
    cin >> K;
    for (int i = 0; i < N; i++) {
        cin >> A[i];
    }

    long long left = -1, right = L + 1;
    while (right - left > 1) {
        long long mid = left + (right - left) / 2;
        if (!check(mid)) {
            right = mid;
        } else {
            left = mid;
        }
    }

    cout << left << endl;
    return 0;
}