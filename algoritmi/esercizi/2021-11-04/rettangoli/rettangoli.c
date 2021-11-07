#include<stdio.h>
#include<math.h>
#include<stdlib.h>

typedef struct {
    float x, y;
} Punto;

typedef struct {
    Punto p1;
    Punto p2;
} Rettangolo;

typedef struct {
    Punto centro;
    Punto fr;
} Cerchio;

Rettangolo *crea_rettangolo() {
    Rettangolo *r;

    r = malloc(sizeof(Rettangolo));
    if(r == NULL) {
        printf("Errore.\n");
        exit(EXIT_FAILURE);
    }

    printf("Inserisci coordinate del primo punto\n");
    scanf("%f%f", &r -> p1.x, &r -> p1.y);
    printf("Inserisci coordinate del secondo punto\n");
    scanf("%f%f", &r -> p2.x, &r -> p2.y);
    return r;

}

int main(){
    float b, h, area_r, duep_r;
   
    printf("RETTANGOLO:\n");

    Rettangolo *ret;
    ret = crea_rettangolo();

    b = (ret -> p2.x) - (ret -> p1.x);
    h = (ret -> p2.y) - (ret -> p1.y);

    area_r = b * h;
    duep_r = 2 * (b + h);

    printf("L'area del rettangolo vale %f, il perimetro vale %f\n", area_r, duep_r);
   return 0;
}
