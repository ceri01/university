#include <stdio.h>
#include <stdlib.h>
#define START 0

char **create_matrix(int n);

int crystal_side(int t);

void print_matrix(char **m, int n);

void crystal(char **m, int r0, int c0, int l);

int main() {
    int t, l;
    char **matrix;

    printf("Inserisci tempo: ");
    scanf("%d", &t);
    
    l = crystal_side(t);
    printf("%d\n", l);
    matrix = create_matrix(l);
    crystal(matrix, START, START, l);
    print_matrix(matrix, l);
}

/*
 *  This function generates a crystal.
 *  Side effects: matrix can be modified.
 *  Post-condition: The matrix will be filled with crystal.
 */

void crystal(char **m, int r0, int c0, int l) {
    m[r0+(l/2)][c0+(l/2)] = '*';
    if (l == 1) {
        m[r0][c0] = '*';
    } else {
        crystal(m, r0, c0, l/2);
        crystal(m, r0, c0+(l/2)+1, l/2);
        crystal(m, r0+(l/2)+1, c0, l/2);
        crystal(m, r0+(l/2)+1, c0+(l/2)+1, l/2);
    }
}

/*
 *  This function finds the value of crystal's side.
 *
 *  Pre-condition: n >= 0 .
 *  Post-condition: return crystal size..
 */

int crystal_side(int t) {
    if (t == 0) {
        return 1;
    } else {
        return (1 + (2 * crystal_side(t-1)));
    }
}

/*
 *  This function creates an empty matrix n X n (empty = '.').
 *
 *  Pre-condition: n > 0 .
 *  Side effects: standard output can be modified.
 *  Post-condition: address of matrix is returned.
 */

char **create_matrix(int n) {
    char **m;

    m = malloc(n * sizeof(char *));
    if(m == NULL) {
        printf("Errore durante l'allocazione.\n");
        exit(EXIT_FAILURE);
    }

    for(int i = 0; i < n; i++) {
        m[i] = malloc(n * sizeof(char));
        if(m == NULL) {
            printf("Errore durante l'allocazione.\n");
            exit(EXIT_FAILURE);
        }
        for (int k = 0; k < n; k++) {
            m[i][k] = '.';
        }
    }
    
    return m;
}

/*
 *  This function prints matrix n X n
 *
 *  Pre-condition: n > 0 && m != NULL.
 *  Side effects: standard output can be modified..
 */

void print_matrix(char **m, int n) {
    for (int i = 0; i < n; i++) {
        for (int k = 0; k < n; k++) {
            printf("%c ", m[i][k]);
        }
        printf("\n");
    }
}


