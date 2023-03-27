#include <stdio.h>

void max_secondmax(int a[], int n, int *max, int *second_max);

int main() {
    int dim, el;
    int max, second_max;
    
    printf("Inserisci lunghezza array\n");
    scanf("%d", &dim);

    int arr[dim]; 
   
    if (dim <= 0) {
        printf("Dimensione non corretta\n");
        return 0;
    }

    printf("Inserisci elementi:\n");
    for(int i = 0; i < dim; i++) {
        scanf("%d", &el);
        arr[i] = el;
    }
    if(dim == 1) {
        printf("max e second maxc coincidono (valore = %d)", arr[0]);
        return 0;
    }
    max_secondmax(arr, dim, &max, &second_max);
    printf("max = %d\nsecond max = %d\n", max, second_max);
}

void max_secondmax(int a[], int n, int *max, int *second_max) {
    int *p = a;
    int tmp;
    
    if(*p > *(p + 1)) {
        *max = *p;
        *second_max = *(p + 1);  
        p = p + 2; 
    } else {
        *max = *(p + 1);
        *second_max = *p;
        p = p + 2; 
    }

    while(p < a + n) {
        if(*p > *max) {
            tmp = *max;        
            *max = *p;
            *second_max = tmp;
        } else if(*p > *second_max) { 
            *second_max = *p;
        }
        p++;
    }
}
