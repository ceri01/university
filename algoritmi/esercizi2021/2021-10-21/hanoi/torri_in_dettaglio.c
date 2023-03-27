#include <stdio.h>
#define ARRS 3
#define ARR 50

void hanoi(int p[ARRS][ARR], int n, int from, int tmp, int to);

void fill(int n, int p[ARRS][n]);

void move_elements(int p[ARRS][ARR], int n, int f, int t); 

int main() {
    int n;
    int pile[ARRS][ARR];

    printf("Inserisci altezza pila: ");
    scanf(" %d", &n);
    printf("\n");

    fill(n, ARRS, pile[ARRS][n]);
    hanoi(pile[ARRS][n], n, n, 0, 0);
}

void fill(int n, int  int p[ARRS][n]) {
    int l = 65;
    printf("kek");
    for (int i = 0; i < ARRS; i++) {
        for(int j = 0; j < n; j++) {
            printf("kek");
            if (i == ARRS) {
                p[ARRS][j] = l;
                l++;
            } else {
                p[ARRS][j] = 0;
            }
        }   
        printf("\n");
    }
}

void move_elements(int p[ARRS][ARR], int n, int f, int t) {
    int tmp = 0;
    // printf("kek");   
    for(int i = 0; i < n; i++) {
        if (p[f][i] == 0) {
            tmp = i-1;
        }
    }
    for(int i = 0; i < n; i++) {
        if(p[t][i] == 0) {
            p[t][i] = p[f][tmp];
            p[f][tmp] = 0;
        }
    }
}

void hanoi(int p[ARRS][ARR], int n, int from, int tmp, int to) {
    if (n == 1) {
        // printf("kek");
        move_elements(p, n, from, to);
        return;
    }
    hanoi(p[ARRS][n], n-1, from, to, tmp);
    move_elements(p, n, from, to); 
    hanoi(p[ARRS][n], n-1, tmp, from, to);
}
