#include <stdio.h>

void hanoi(int n, int from, int tmp, int to, int * c);

int main() {
    int n, count = 0;
    printf("Inserisci altezza pila: ");
    scanf(" %d", &n);
    hanoi(n, 0, 1, 2, &count);
    printf("\n%d", count);
}

void hanoi(int n, int from, int tmp, int to, int * c) {
    if (n == 1) {
        *c = *c + 1;
        return;
    }
    hanoi(n-1, from, to, tmp, c);
    *c = *c + 1;
    hanoi(n-1, tmp, from, to, c);
}
