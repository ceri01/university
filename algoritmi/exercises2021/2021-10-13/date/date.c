#include <stdio.h>
#define L 100

typedef struct {
    int g;
    int m;
    int a;
}  Data;

int main() {
    Data date[L];
    Data d;
    int index = 0;
    printf("Inserisci le date\n");
    for (; index < L - 2; index++) {
        scanf("%d/%d/%d", &d.g, &d.m, &d.a);
        if (d.a == 0 && d.m == 0 && d.g == 0) {
            break; 
        }
        if (d.a < 0 || (d.m > 12 || d.m < 0) || (d.g > 31 || d.g < 0)) {
            printf("Errore inserimento dati\n");
            return 0;
        }
        date[index] = d;
    } 

    printf("Inserisci data per confronto:\n");
    scanf("%d/%d/%d", &d.g, &d.m, &d.a);
    printf("date precedenti all'ultima data inserita:\n");
    for (int i = 0; i < index; i++) {
        if ((date[i].a < d.a) && (date[i].m < d.m) && (date[i].g < d.g))
            printf("%02d/%02d/%04d\n", date[i].g, date[i].m, date[i].a); 
    } 

}
