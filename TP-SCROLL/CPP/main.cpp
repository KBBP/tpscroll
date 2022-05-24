#include<iostream>
using namespace std;

int main() {
    int grade[] = { 3, 1, 6, 30, 8, 9, 2, 5 };
    int length = sizeof(grade) / sizeof(*grade);
    int temp;

    for (int a = 1; a < length; a++) {
        for (int b = 0; b < length - a; b++) {
            if (grade[b] > grade[b + 1]) {
                temp = grade[b];
                grade[b] = grade[b + 1];
                grade[b + 1] = temp;
            }
        }
    }

    cout << "Hasil sorting: " << endl;
    for (int x = 0; x < length; x++) {
        cout << " " << grade[x];
    }
}