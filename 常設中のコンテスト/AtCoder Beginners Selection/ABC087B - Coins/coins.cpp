#include <iostream>
using namespace std;

int A, B, C, X;

int main() {
    cin >> A;
    cin >> B;
    cin >> C;
    cin >> X;

    int ans = 0;
    int i, j, k;
    for (i = 0; i <= min(X/500, A); i++) {
        for (j = 0; j <= min((X-i*500)/100, B); j++) {
            int rest = X - i*500 - j*100;
            if (rest >= 0 && rest % 50 == 0 && rest / 50 <= C) {
                ans++;
            }
        }
    }

    cout << ans << endl;
    return 0;
}
