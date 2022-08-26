#include <iostream>
using namespace std;

int main() {
    int n;
    cin >> n;

    int sumA = 0;
    int sumB = 0;
    for (int i = 0; i < n; i++) {
        int a, b;
        cin >> a >> b;
        
        if (a > b) {
            sumA += a + b;
        } else if (a < b) {
            sumB += a + b;
        } else {
            sumA += a;
            sumB += b;
        }
    }

    cout << sumA << " " << sumB << endl;

    return 0;
}
