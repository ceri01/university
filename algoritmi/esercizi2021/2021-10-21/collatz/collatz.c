#include <stdio.h>

int next_in_collatz(int n);

int main() {
    int num, lung;
    printf("Inserisci un numero: ");
    scanf(" %d", &num);
    printf("%d ", num);
    
    for (lung = 1; num != 1; lung++) {
        if (num != 1) {
            num = next_in_collatz(num);
            printf("%d ", num);
        }
    }
    printf("\nLunghezza: %d\n", lung);
}

int next_in_collatz(int n) {
    if(n % 2 == 0) {
        n = n / 2;
    } else {
        n = (n * 3) + 1;
    }
    return n;
}
