
// MOOOOLTO inefficiente!!


#include <stdio.h>
#include <time.h>
#include <math.h>
#include <stdlib.h>
#include <stdbool.h>
#define DIM 10
#define DOT 46
#define CHARS 26

void print_matrix(char m[DIM][DIM]) {
    for(int i = 0; i < DIM; i++) {
        for(int k = 0; k < DIM; k++) {
            printf("%c", m[i][k]);
        }
        printf("\n");
    }
}

void fill_matrix(char m[DIM][DIM]) {
    for(int i = 0; i < DIM; i++) {
        for(int k = 0; k < DIM; k++) {
            m[i][k] = '.';
        }
    }
}

bool is_free(char m[DIM][DIM], int main_coor, int x, int y) {
    if (m[y][x] != DOT || main_coor > 9 || main_coor < 0) {
        return false;
    }
    return true;
}

int check_around(char m[DIM][DIM], int x, int y) {
    int pos[4] = {-1, -1, -1, -1}, i = 0, r;
    if(is_free(m, y-1, x, y-1)) {
        pos[i] = 0;
        i++;
    }
    if(is_free(m, x+1, x+1, y)) {
        pos[i] = 1;
        i++;
    }
    if(is_free(m, y+1, x, y+1)) {
        pos[i] = 2;
        i++;
    }
    if(is_free(m, x-1, x-1, y)) {
        pos[i] = 3;
        i++;
    }
    if (pos[0] == -1) {
        return -1;
    }
    r = rand() % i;
    return pos[r];
}


int main() {
    char mat[DIM][DIM];
    char c = 'A';

    fill_matrix(mat);
    srand(time(NULL));

    int x = rand() % 10; // scorre orizzontalmente (secomdo indice mat)
    int y = rand() % 10; // scorre verticalmente (primo indice mat)
    int dir;
    bool done = true;
    
    mat[y][x] = c; // Partenza
    c++; 
    for(int index = 1; index < CHARS; index++) {
        if(done) {
            dir = rand() % 4;
        }
        done = false;
        switch(dir) {
            case 0: //su
                y--;
                if(is_free(mat, y, x, y)) {
                    mat[y][x] = c;
                    done = true;
                    c++;
                    break;
                }
                y++;
                break;
            case 1: // dx
                x++;
                if(is_free(mat, x, x, y)) {
                    mat[y][x] = c;
                    done = true;
                    c++;
                    break;
                }
                x--;
                break;
            case 2: // giu
                y++;
                if(is_free(mat, y, x, y)) {
                    mat[y][x] = c;
                    done = true;
                    c++;
                    break;
                }
                y--;
                break;
            case 3: // sx
                x--;
                if(is_free(mat, x, x, y)) {
                    mat[y][x] = c;
                    done = true;
                    c++;
                    break;
                }
                x++;
                break;
        }
        if (!done) {
            dir = check_around(mat, x, y);
            if (dir == -1) {
                break;
            }
            index--;
            continue;
        }
    }
    print_matrix(mat);
}
