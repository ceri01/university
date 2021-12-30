#include <stdio.h>

int main() {
    int r = 0, c = 0;

    printf("Inserisci primo indice: ");
    scanf(" %d", &r);
    printf("Inserisci secondo indice: ");
    scanf(" %d", &c);

    char mat[r][c], ch;
    
    for(int i = 0; i < r; i++) {
        for(int k = 0; k < c; k++) {
            scanf(" %c", &ch);
            mat[i][k] = ch;
        }
    }

    for(int i = r; i > 0; i--) {
        for(int k = c; k > 0; k--) {
            if (mat[i][k] != '*') {
                for(int j = k;mat[i][k] == '*';) { 
                    mat[i][k] = mat[i-1][k];
                    mat[i-1][k] = '*';   
                }
            }
        }
    }
    
    for(int i = 0; i < r; i++) {
        for(int k = 0; k < c; k++) {
            printf("%c ", mat[i][k]);
        }
        printf("\n");
    }
}
