#include <stdio.h>
#include <string.h>

void switch_pos( char *a[], int n, int i ) {
    if ( i < 0 || i > n-1 )
        return;
        
    char *p = a[i];
    a[i] = a[n-1];
    a[n-1] = p;
    return;
}

int main( int argc, char **argv ) {
    char **q;
    int i;
    
    scanf( "%d", &i );
    switch_pos( argv + 1, argc - 1, i );

    for ( q = argv + 1; q < argv + argc; q++ )
        printf( "%s ", *q );
    printf( "\n" );
    return 0;
}
