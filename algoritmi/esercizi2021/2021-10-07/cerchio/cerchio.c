#include <stdio.h>
#include <math.h>

float area(float raggio) {
    return 2*raggio*M_PI;
}

void main() {
    float r;
    float ar;
    printf("Inserisci il raggio del cerchio \n");
    scanf(" %f", &r);
    ar = area(r);
    printf("L'area del cerchio con raggio %f è %f\n", r, ar);
}

