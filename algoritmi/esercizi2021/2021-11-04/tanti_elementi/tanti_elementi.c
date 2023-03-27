#include <stdio.h>
#include <stdlib.h>

void *my_malloc(size_t size) {
    void *p;
    p = malloc(size);
    if(p == NULL) {
        printf("Errore allocazione\n");
        exit(EXIT_FAILURE);
    }
    return p;
}

char *read_n( int *num ) {
    char *a;

    scanf( "%d", num );
    a = my_malloc( *num * sizeof(char) );
    
    for (int i = 0; i < *num; i++ ) {
        scanf(" %c", (a + i));
    }
    printf("\n");

    return a;
}

int main() {
    char *p; 
    int *n;
    p = read_n(n);
    for(char *ptr = p; ptr < p + *n; ptr++) {
        printf("%c ", *ptr);
    }
    printf("\n");
}
