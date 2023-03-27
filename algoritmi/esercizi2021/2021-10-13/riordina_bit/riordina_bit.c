#include <stdio.h>
#include <stdlib.h>

void sort(int * a, int dim);    

int main() {
    int n, el;

    printf("Inserisci il numero di bit: ");
    scanf(" %d", &n);

    int arr[n];

    for(int i = 0; i < n; i++) {
        scanf(" %d", &el);
        arr[i] = el;
    }

    sort(arr, n);

    for(int i = 0; i < n; i++) {
        printf("%d ", arr[i]);
    }
}

void sort(int * a, int dim) {
    int init = 0, end = dim - 1;
    int tmp;

    for (;init < end;) { 
        for (;a[init] == 0;) {
            init++;
        }
    
        for (;a[end] == 1;) {
            end--;
        }
        
        if (init >= end) {
            break;
        }
    
        tmp = a[init];
        a[init] = a[end];
        a[end] = tmp;
    }
}
