#include <stdio.h>

void hanoi(int n, int from, int tmp, int to);

int main() {
    int n;
    printf("Inserisci altezza pila: ");
    scanf(" %d", &n);
    printf("\n");

    hanoi(n, 0, 1, 2);
}

void hanoi(int n, int from, int tmp, int to) {
    if (n == 1) {
        printf("%d -> %d\n", from, to);
        return;
    }
    hanoi(n-1, from, to, tmp);
    printf("%d -> %d\n",from, to);
    hanoi(n-1, tmp, from, to);
}
