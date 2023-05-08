#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>

#define SIZE 5

char *read_line(char c);
char *read_word(void);

int main() {
    char c;
    char *s1, *s2;
    
    printf("Inserisci la lettea che interrompera' l'acquisizione della parola:\n");
    scanf(" %c", &c);
    s1 = read_line(c);
    
    printf("Inserisci una parola (se contiene caratteri non alfanumerici si interrompera' l'acquisizione)\n");
    s2 = read_word();
    
    
    if (s1 == NULL || s2 == NULL) {
        printf("Si e' verificato un errore.\n");
    }
    
    printf("%s --- %s\n", s1, s2);
    
    free(s1);
    free(s2);
}

char *read_word(void) {
    char ch;
    char *str;
    int size = SIZE, i = 0;

    str = malloc(size * sizeof(char));
    
    if(str == NULL) {
        printf("Errore allocazione\n");
        return str;
    }
    while((ch = getchar()) != ' ') { 
        if(!isalnum(ch)) {
            break;
        }
        if(i >= size) {
            size += 5;
            str = realloc(str, size);
        }
        *(str + i) = ch;
        i++;
    }
    
    if(i >= size) {
        str = realloc(str, (size + 1));
    } 
    
    *(str + i) = '\0';
    getchar();
    return str;
}

char *read_line(char c) {
    char ch;
    char *str;
    int size = SIZE, i = 0;

    str = malloc(size * sizeof(char));
    
    if(str == NULL) {
        printf("Errore allocazione\n");
        return str;
    }


    while((ch = getchar()) != c) {
        if(i >= size) {
            size += 5;
            str = realloc(str, size);
        }
        *(str + i) = ch;
        i++;
    }
    
    if(i >= size) {
        str = realloc(str, (size + 1));
    } 
    
    *(str + i) = '\0';
    getchar();
    return str;
}
