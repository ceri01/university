#include <stdio.h>

int main() {
    int n, count = 0;
    printf("Inserisci un nuermo: ");
    scanf(" %d", &n);

    for (int i = n; i != 0; i--) {
        if ((n % i) == 0) {
            count++;
            printf("divisore di %d = %d\n", n, i);
        }
    }
    printf("Numero di divisori: %d\n", count);
}
