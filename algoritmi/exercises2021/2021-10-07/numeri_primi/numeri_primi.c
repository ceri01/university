#include <stdio.h>

int main() {
    int n, count = 0;
    printf("Inserisci un nuermo: ");
    scanf(" %d", &n);

    for (int i = n; i > 0; i--) {
        if ((n % i) == 0) {
            count++;
        }
    }

    if (count == 2) {
        printf("%d è un numero primo\n", n);
    } else {
        printf("%d non è un numero primo\n", n);
    }
}
