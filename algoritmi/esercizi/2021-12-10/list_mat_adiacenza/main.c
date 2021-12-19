#include <stdio.h>
#include "code_liste_bidirezionali/b_list.h"
#include "array_di_liste.h"

int main() {
    int n = 8;
    int **m = new_matr(n);
    insert(m, 0, 2);
    insert(m, 0, 5);
    insert(m, 2, 1);
    insert(m, 2, 5);
    insert(m, 3, 2);
    insert(m, 4, 5);
    insert(m, 5, 6);
    insert(m, 6, 4);
    insert(m, 7, 6);
    print_mat(m, n);
    matr2arrayDiListe(m, n);
}
