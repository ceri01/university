#include <stdio.h>

int binary_search(int * a, int l) {
    int init = 0, end = l;
    if (end == a[end]) {
        return end + 1;
    }
    int m = (end + init) / 2;
    for(;init != end;) {
        if(a[m] == m) {
            init = m + 1;
            m = (end + init) / 2;
        } else {
            end = m;
            m = (end + init) / 2;
        }
    }
    return init;
}


int main() {
    int n, last = 0, index;
    printf("Inserisci il numero di elementi che contiene il vettore: ");
    scanf("%d", &last);
    
    int arr[last];
    if (last == 0) {
        printf("l'unico elemento inserito e' lo 0\n");
        return 0;
    }
    for(index = 0; index < last; index++) {
        printf("Valore inserito: ");
        scanf(" %d", &n);
        arr[index] = n;
    }

    printf("elemento mancante: %d\n", binary_search(arr, index));
}
