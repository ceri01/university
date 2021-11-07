#include <stdlib.h>
#include <stdio.h>

void *my_malloc( size_t n ) {
    void *p;
    p = malloc( n );
    if ( p == NULL ) {
        printf( "Allocazione fallita, p punta a NULL.\n" );
        exit( EXIT_FAILURE );
    }
    return p;
}

void *my_realloc(void *p, size_t size) {
    p = realloc(p, size);
    if (p == NULL) {
        printf("Riallocazione fallita, p punta a NULL.\n");
        exit(EXIT_FAILURE)j;
    }
    return p;
}
