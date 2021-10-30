#include <stdio.h>
#include <string.h>

typedef char* String;

int smallest_word_index(char *s[], int n);

String *smallest_word_indexx(String s[], int n);

int main( void ) {
    String dict[] = { "ciao", "mondo", "come", "funziona", "bene", "il", "programma" };
    int lun = 7;
    printf("%d\n", smallest_word_index(dict, lun));
    printf("%s\n", *smallest_word_indexx(dict, lun));
}

int smallest_word_index(String s[], int n) {
    String min = s[0];
    int index = 0;

    for(int i = 1; i < n; i++) {
        if(strcmp(min, s[i]) > 0) {
            min = s[i];
            index = i;
        }
    }

    return index;
}

String *smallest_word_indexx(String s[], int n) {
    String min = s[0];

    for(int i = 1; i < n; i++) {
        if(strcmp(min, s[i]) > 0) {
            min = s[i];
        }
    }
    printf("%s", min);
    printf("%p", &min);   
    return min;
}
