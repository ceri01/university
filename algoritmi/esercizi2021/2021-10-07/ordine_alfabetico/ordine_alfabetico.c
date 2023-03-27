#include <stdio.h>
#include <math.h>

int main() {
    char car_1, car_2;
    int dist;

    for(;;) {
        printf("Inserisci primo carattere (Maiuscolo): ");
        scanf(" %c", &car_1);
        printf("Inserisci secondo carattere (Maiuscolo): ");
        scanf(" %c", &car_2);
        if ((car_1 <= 90 && car_1 >= 65) && (car_2 <= 90 && car_2 >= 65)) {
            break;
        }
    }

    dist = abs(car_1 - car_2) + 1;
    printf("\nLa distanza tra %c e %c è %d\n", car_1, car_2, dist);
}
