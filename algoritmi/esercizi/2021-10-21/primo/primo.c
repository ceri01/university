#include <stdio.h>
#include <stdbool.h>

bool primo(int n);

int main() {
    int n;
    printf("Inserisci un nuermo: ");
    scanf(" %d", &n);

    if (primo(n)) {
        printf("%d è un numero primo\n", n);
    } else {
        printf("%d non è un numero primo\n", n);
    }
}

bool primo(int n) {
    for(int i = 2; i < n; i++) {
        if (n % i == 0) {
            return false;
        }
    }
    return true;
}
