#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#include <string.h>

#define N 100

void treVolte( char *a[], int n ) {
    char **p, *q = a[0];
    int conta = 0;

    for ( p = a; p < a + n; p++ ) {
        int contae = 0;
        char *c;
        c = *p;
 
        while ( *c != ' ' ) {
            printf("%c", *c);
            if ( *c == 'e' )
                contae++;
            if ( contae >= 3 ) {
                conta++;
                break;
            }
            c++;
        }
       // printf("%d\n", conta);
        
        if ( conta == 3 ) {
            strcpy( q, *p );
            strcpy( *p, a[0]);
            strcpy ( a[0], q);
            return;
        }
    }
}

int main( int argc, char **argv ) {
    treVolte( argv + 1, argc - 1 );
    
    for ( char** p = argv + 1; p < argv + argc; p++ )
        printf( "%s ", *p );
    printf( "\n" );
    return 0;
}
