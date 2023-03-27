#include <stdio.h>
#include <math.h>
#include <stdlib.h>

int main() {
    int a, b, c;
    float delta;
    float p_int, p_imm = 0;

    printf("Inserisci il valore di a: ");
    scanf(" %d", &a);
    printf("Inserisci il valore di b: ");
    scanf(" %d", &b);
    printf("Inserisci il valore di c: ");
    scanf(" %d", &c);
    
    delta = pow(b, 2) - (4*a*c);
    p_int = (-b)/(2*a);
    if (delta < 0) {
        p_imm = sqrt(abs(delta))/(2*a);
        printf("x1 = %.2f + i* %.2f\n", p_int, p_imm);
        printf("x2 = %.2f - i* %.2f\n", p_int, p_imm);
    } else {
        float tmp = sqrt(delta)/(2*a);
        printf("x1 = %.2f\n", p_int + tmp);
        printf("x2 = %.2f\n", p_int - tmp);
    }
}
