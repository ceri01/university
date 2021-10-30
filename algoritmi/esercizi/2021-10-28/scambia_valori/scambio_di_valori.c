#include <stdio.h>

void scambia(int *p, int *q);

int main() {
    int p, q;
    scanf("%d %d", &p, &q);
    scambia(&p, &q);
    printf("%d %d\n", p, q);
}

void scambia(int *p, int *q) {
    int tmp = *p;
    *p = *q;
    *q = tmp;
}


