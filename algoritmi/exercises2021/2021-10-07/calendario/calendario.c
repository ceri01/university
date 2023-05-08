#include <stdio.h>


int get_days_in_month(int m) {
    switch(m) {
        case 1: case 3: case 5: case 7: case 8: case 10: case 12:
            return 31;
        case 2:
            return 28;
        default:
            return 30;
    }
}




int main() {
    int month, init_day;

    printf("Inserisci il numero del mese (1 = gennaio, ..., 12 = dicembre): ");
    scanf(" %d", &month);
    printf("Inserisci il giorno di inizio mese (1 = lunedi, ···, 7 = domenica): ");
    scanf(" %d", &init_day);

    if (month > 12 || month < 1 || init_day > 7 || init_day < 1) {
        printf("Errore di inserimento");
        return 0;
    }

    int d_in_m = get_days_in_month(month);

    printf(" L   M   M   G   V   S   D\n");

    for (int i = 0, day = 1; (i <= 5); i++) {
        for (int k = 0; k < 7 && d_in_m >= day; k++) {
            if (i == 0 && k+1 < init_day) {
                printf("    ");
                continue;
            } else {
                if (day < 10) {
                    printf(" %d  ", day);
                } else {
                    printf(" %d ", day);
                }
                day++;
            }
        }
        printf("\n");
    }
}
