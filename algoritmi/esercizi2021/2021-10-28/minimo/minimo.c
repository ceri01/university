#include <stdio.h>

int *smallest(int a[], int n);

int main() {
    int len;
    int elem;
    printf("Inserisci lunghezza vettore\n");
    scanf(" %d", &len);

    if (len <= 0) {
        printf("Errore! l'array deve contenere almeno un elemento\n");
        return 0;
    }

    int arr[len];
    printf("Inserisci %d valori\n", len);
    
    for(int *i = arr; i < arr + len; i++) {
        scanf("%d", &elem);
        *i = elem;
    }
    printf("Valore minimo: %d\n", *smallest(arr, len));
}

int *smallest(int arr[], int n) {
    int *min = arr;

    for(int *ptr = arr + 1; ptr < arr + n; ptr++) {
        if (*ptr <= *min) {
            min = ptr;
        }
    }
    return min;
}
