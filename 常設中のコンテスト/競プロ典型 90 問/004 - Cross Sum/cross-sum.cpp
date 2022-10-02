#include <iostream>
using namespace std;

int H, W;

int main() {
    cin >> H >> W;

    int A[H][W];
    for (int i = 0; i < H; i++) {
        for (int j = 0; j < W; j++) {
            cin >> A[i][j];
        }
    }

    int row[H];
    for (int i = 0; i < H; i++) {
        row[i] = 0;
        for (int j = 0; j < W; j++) {
            row[i] += A[i][j];
        }
    }

    int col[W];
    for (int i = 0; i < W; i++) {
        col[i] = 0;
        for (int j = 0; j < H; j++) {
            col[i] += A[j][i];
        }
    }

    for (int i = 0; i < H; i++) {
        for (int j = 0; j < W; j++) {
            cout << row[i] + col[j] - A[i][j] << " ";
        }
        cout << endl;
    }
}