#include <stdio.h>

int potenza(int b, int e); 

int main() {
    int numero, esponente;

    printf("Inserisci un numero: ");
    scanf(" %d", &numero);
    printf("Inserisci un numero che verra' usato come esponente: ");
    scanf(" %d", &esponente);
    printf("%d alla %d fa %d", numero, esponente, potenza(numero, esponente));
}

int potenza(int b, int e) {
    if (e == 1) {
        return b;
    } else {
        return b * potenza(b, e - 1);
    }
}
