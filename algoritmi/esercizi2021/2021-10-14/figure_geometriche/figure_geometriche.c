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

int main(){
    float b, h, raggio, area_r, duep_r, area_c, circ_c;
    Rettangolo r;
    Cerchio c;

    printf("RETTANGOLO:\n");
    printf("Inserisci le coordinate del punto in basso a sinistra\n");
    scanf("%f%f", &r.p1.x, &r.p1.y);
    printf("Inserisci le coordinate del punto in alto a destra\n");
    scanf("%f%f", &r.p2.x, &r.p2.y);
    printf("CERCHIO:\n");
    printf("Inserisci le coordinate del centro del cerchio\n");
    scanf("%f%f", &c.centro.x, &c.centro.y);
    printf("Inserisci le coordinate del punto finale del raggio del cerchio\n");
    scanf("%f%f", &c.fr.x, &c.fr.y);
    b = r.p2.x - r.p1.x;
    h = r.p2.y - r.p1.y;
    raggio = sqrt(pow(abs(c.centro.x - c.fr.x), 2) + pow(abs(c.centro.y - c.fr.y), 2));
    area_r = b * h;
    duep_r = 2 * (b + h);
    area_c = M_PI * (raggio * raggio);
    circ_c = 2 * M_PI * raggio;
    printf("L'area del rettangolo vale %f, il perimetro vale %f\n", area_r, duep_r);
    printf("L'area del cerchio vale %f, la circonferenza vale %f\n", area_c, circ_c);
    return 0;
}
