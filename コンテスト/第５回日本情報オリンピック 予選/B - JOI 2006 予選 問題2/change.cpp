#include <iostream>
#include <map>
using namespace std;

int main() {
    int n;
    cin >> n;

    map<char, char> dict;
    for (int i = 0; i < n; i++) {
        char key, value;
        cin >> key >> value;
        dict[key] = value;
    }

    int m;
    cin >> m;
    
    for (int i = 0; i < m; i++) {
        char key;
        cin >> key;
        if (dict.find(key) != dict.end()) {
            cout << dict[key];
        } else {
            cout << key;
        }
    }
    cout << endl;

    return 0;
}
