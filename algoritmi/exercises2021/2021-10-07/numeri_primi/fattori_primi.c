#include <stdio.h>
#include <math.h>

int main() {
    int n, tmp;
    printf("Inserisci un nuermo: ");
    scanf(" %d", &n);

    tmp = n;

    for (int i = 2; i != n;) {
        if ((tmp % i) == 0) {
            tmp = tmp / i;
            printf("%d ", i);
        } else {
            i++;
        }
    }
}
