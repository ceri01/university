#ifndef ARRAY_DI_LISTE_H
#define ARRAY_DI_LISTE_H
#include "code_liste_bidirezionali/b_list.h"

typedef struct list_node **arrayDiListe;

// Trasforma una matrice di adiacenza in un array di liste di adiacenza
arrayDiListe matr2arrayDiListe(int **matr, int n);

// Crea una nuova matrice di adiacenza
int **new_matr(int n);

// Stampa la matrice di adiacenza
void print_mat(int **matr, int n);

// Inserisce un arco nel grafo
void insert(int **matr, int n1, int n2);

// kek ma anche w
int **f_matr(int **matr, int n);

#endif
