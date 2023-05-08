// 1) non ènecessario memorizzare tutta la frase
//
// 2) ènecessario memorizzare il numero di occorrenze delle lettere

#include <stdio.h>
#include <ctype.h>

int main() {
    char c,  character[26] = {0};
    
    // 2n confronti
    while((c = getchar()) != '.') {
        if (isalpha(c)) {
            c = tolower(c);
            character[c - 97]++;
        }
    }
    // ciclo eseguito 26 volte
    for (int i = 0; i < 26; i++) {
        // 26 confronti
        if (character[i] > 0) {
            printf("%c ", (i+97));
            // eseguito al più 26 vole con un totale di circa n iterazioni
            for(int j = character[i]; j > 0; j--) {
                printf("* ");
            }
            printf("\n");
        }
    } 
}

// 3 variabili e un array da 26 elementi

