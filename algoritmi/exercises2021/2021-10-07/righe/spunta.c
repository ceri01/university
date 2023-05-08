#include <stdio.h>
#include <math.h>

int main(void) {
    int n, i, j;
    scanf("%d", &n);
    for (i = 0; i < n; i++) {
        for (j = n+1; j >= 0; j--) {
            if (((n - 3) == i && j == n + 1) || i == (n - 2) && j == n || j == i) {
                printf("*");
            } else {
                printf(".");
            }
        }
        printf("\n");
    }
}
