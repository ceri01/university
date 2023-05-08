#include <stdio.h>

int main() {
    int n, i, count = 0;
    printf("Inserisci un nuermo: ");
    scanf(" %d", &n);
    i = n;

    while (i > 0) {
        if ((n % i) == 0) {
            count++;
        }
        i--;
    }

    if (count == 2) {
        printf("%d è un numero primo\n", n);
    } else {
        printf("%d non è un numero primo\n", n);
    }
}
