#include <iostream>
using namespace std;

int N;

bool check(string s) {
    int count = 0;
    for (int i = 0; i < N; i++) {
        if (s[i] == '(') {
            count++;
        }
        if (s[i] == ')') {
            count--;
        }
        if (count < 0) {
            return false;
        }
    }
    return count == 0;
}

int main() {
    cin >> N;

    for (int i = 0; i < (1 << N); i++) {
        string s;
        for (int j = N-1; j >= 0; j--) {
            if (i & (1 << j)) {
                s += ')';
            } else {
                s += '(';
            }
        }
        if (check(s)) {
            cout << s << endl;
        }
    }

    return 0;
}
