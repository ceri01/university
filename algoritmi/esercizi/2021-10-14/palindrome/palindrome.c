#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#define vero 1
#define falso 0

typedef int Bool;

char * get_word(int s, int * pt_s) {
    char c, *w = NULL;
    int tmp_size = 0;
    w = malloc(s);
    while((c = getchar()) != '.') {
        if(tmp_size >= s) {
            s = s + 10;
            w = realloc(w, s);
        }
        w[tmp_size++] = c;
    }
    *pt_s = tmp_size;
    return w;
}

int main() {
    char *word;
    Bool pal = vero;
    int size = 5;
    int *pt_size = &size;
    word = get_word(size, pt_size);
    //printf("%s %d", word, size);
    
    for (int i = 0, k = size-1; i < size; i++, k--) {
        if (word[i] != word[k]) {
            pal = falso;
            break;
        }
        if (k == 1) {
            break;
        }
    }

    if (pal) {
        printf("%s è palindroma\n", word);
    } else {
        printf("%s non è palindroma\n", word);
    }
    // free(pt_size);
}
