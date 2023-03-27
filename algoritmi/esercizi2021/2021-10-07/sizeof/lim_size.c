#include <stdio.h>
#include <limits.h>
#include <float.h>
#include <stdlib.h>

int main() {
    int a = 3;
    short b = 4;
    long c = 78;
    char d = 's';
    double e = 2.34;
    float f = 3.33;

    printf("byte di a (int) = %d --> max = %d min == %d", sizeof(a), INT_MAX, INT_MIN);
    printf("\nbyte di b (short) = %d -i-> max = %d min == %d\n", sizeof(b), SHRT_MAX, SHRT_MIN);  
    printf("\nbyte di c (long) = %d --> max = %ld min == %ld\n", sizeof(c), LONG_MAX, LONG_MIN);
    printf("\nbyte di d (char) = %d --> max = %c min == %c\n", sizeof(d), SCHAR_MAX, SCHAR_MIN);
    printf("\nbyte di e (double) = %d--> max = %f min == %f\n", sizeof(e), DBL_MAX, DBL_MIN);
    printf("\nbyte di f (floaT) = %d --> max = %f min == %f\n", sizeof(f), FLT_MAX, FLT_MIN);
_}
