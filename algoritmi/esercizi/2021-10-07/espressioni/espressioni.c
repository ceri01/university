#include <stdio.h>

int main() {
    int count_open = 0, count_close = 0, count_tot = 0, lp = 0;
    char ch;

    while((ch = getchar()) && ch != '.') {
        if(ch == '(') {
            count_open++;
        } else if(ch == ')') {
            count_close++;
        }
        count_tot++;
        if(count_close > count_open && lp == 0) {
            lp = count_tot;
        }

    }

    if(count_open == count_close) {
        printf("La stringa è un'espressione ben parentesizzata\n");
    } else if (count_open > count_close) {
        printf("La stringa non è ben parentesizzata!\n");
        printf("carattere %d: mancano parentesi chiuse alla fine", count_tot);
    } else {
        printf("La stringa non è ben parentesizzata!\n");
        printf("carattere %d: troppe parentesi chiuse", lp);
    }
}
