#include <iostream>
using namespace std;

int N;
int A[200];

int main() {
    cin >> N;
    for (int i = 0; i < N; i++) {
        cin >> A[i];
    }

    int ans = 0;
    while (true) {
        bool ok = true;
        for (int i = 0; i < N; i++) {
            if (A[i] % 2 == 0) {
                A[i] /= 2;
            } else {
                ok = false;
                break;
            }
        }
        if (ok) {
            ans++;
        } else {
            break;
        }
    }

    cout << ans << endl;
    return 0;
}
