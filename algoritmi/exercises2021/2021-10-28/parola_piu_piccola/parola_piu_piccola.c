#include <stdio.h>
#include <string.h>

typedef char* String;

int smallest_word_index(char *s[], int n);

String *smallest_word_address(String s[], int n);

int main( void ) {
    String dict[] = { "ciao", "mondo", "come", "funziona", "bene", "il", "programma" };
    int lun = 7;
    printf("%d\n", smallest_word_index(dict, lun));
    printf("%s\n", *smallest_word_address(dict, lun));
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

char **smallest_word_address(char * s[], int n) {
    int min_pos = 0;

    for(int i = 1; i < n; i++) {
        if(strcmp(s[min_pos], s[i]) > 0) {
            min_pos = i;
        }
    } 
    return s + min_pos;
    // String s[] = **s = puntatore --> (array di puntatori in cui ogni puntatore punta a un array di caratteri) puntatore --> primo carattere della stringa puntata
    // s = il suo valore e' l'indirizzo del puntatore che punta al primo array di caratteri
    // min_pos posizione della stringa piu' piccola
    // s + min_pos = indirizzo del puntatore che punta alla stringa pi' piccola
}
