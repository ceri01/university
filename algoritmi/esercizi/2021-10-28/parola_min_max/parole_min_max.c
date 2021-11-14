#include <stdio.h>
#include <string.h>

void smallest_larger(char *s[], int n, char **smallest, char **largest);

int main( void ) {
    char *dict[] = { "zzao", "mondo", "come", "afunziona", "bene", "il", "programma" };
    char *smallest, *largest;
    int lun = 7;
    smallest_larger(dict, lun, &smallest, &largest);
    printf("min = %s\nmax = %s\n", smallest, largest);
}

void smallest_larger(char *s[], int n, char **smallest, char **largest) { 
    *smallest = s[0];
    *largest = s[0];
    
    for(int i = 1; i < n; i++) { 
        if(strcmp(s[i], *largest) > 0) {
            *largest = s[i];
            continue;
        }
        if(strcmp(s[i], *smallest) < 0) {
            *smallest = s[i];
            continue;
        }
    }
}

