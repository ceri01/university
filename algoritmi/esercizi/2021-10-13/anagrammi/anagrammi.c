#include <stdio.h>
#include <ctype.h>

int main() {
    char c,  str1[26] = {0}, str2[26] = {0};
    
    // 4n confronti
    while((c = getchar()) != '\n') {
        if (isalpha(c)) {
            c = tolower(c);
            str1[c - 97]++;
        }
    }

    while((c = getchar()) != '\n') {
        if (isalpha(c)) {
            c = tolower(c);
            str2[c - 97]++;
        }
    }
    // ciclo eseguito 26 volte
    for (int i = 0; i < 26; i++) {
        if (str1[i] != str2[i]) {
            printf("Le due parole non sono l'una l'anagramma dell'altra\n");
            return 0;
        }
    }
    printf("Le due parole sono l'una l'anagramma dell'altra\n");
}

