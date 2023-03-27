#include <stdio.h>
#include <stdlib.h>
#include "code_liste_bidirezionali/b_list.h"
#include "array_di_liste.h"

struct list_node {
    struct list_node *next;
    int v;
};

int **new_matr(int n) {
    int **mat = calloc(n, sizeof(int*));
    for(int i = 0; i < n; i++) {
        mat[i] = calloc(n, sizeof(int));
    }
    return mat;
}

void print_mat(int **matr, int n) { 
    printf("    ");
    for(int i = 0; i < n; i++) {
        printf("%d ", i);
    }
    printf("\n   ");
    for(int i = 0; i < n; i++) {
        printf(" _");
    }
    printf("\n");
    for(int i = 0; i < n; i++) {
        for(int k = 0; k < n; k++) {
            if(k == 0) {
                printf("%d | ", i);
            }
            printf("%d ", matr[i][k]);
        }
        printf("\n");
    }   
}

void insert(int **matr, int n1, int n2) {
    matr[n1][n2] = 1;
}

int **f_matr( int **matr, int n ) {
    int **matr2 = new_matr(n);
    for (int i = 0; i < n; i++)
        for (int j = 0; j < n; j++)
            if (matr[i][j] == 1) 
                for (int k = 0; k < n; k++)
                    if (matr[j][k] == 1)
                        matr2[i][k] = 1;  
    return matr2;
}

arrayDiListe matr2arrayDiListe(int **matr, int n) {
    arrayDiListe m = calloc(n, sizeof(struct list_node));
    struct list_node *tmp;
    for(int i = 0; i < n; i++) {
        struct list_node *node;
        for(int j = n; j > n; j--) {
            struct list_node *new;
            new -> v = matr[i][j];
            if(j == n) {
                new -> next = NULL;
            } else {
                new -> next = tmp;   
            }
            tmp = new;
        }
        m[i] = node;
    }
    return m;
}


