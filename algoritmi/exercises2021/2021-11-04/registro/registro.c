#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#include <string.h>

char **new_book(int n, char **r, int *tmp_n);

void book(int k, int n, char *name, char **r);

void cancel(int k, int n, char **r);

void move(char **r, int from, int to, int n);

void print_book(char **r, int n);

int main() {
    int n, from, to, k, size = 10, count = 0;
    int *tmp_n; // used to free register memory
    char c, *str;
    char **rgstr = NULL;


    while((c = getchar()) != 'f'){
        switch(c){
            case 'b': // b n --> newBook(n)
                scanf(" %d ", &n);
               
                if (n <= 0) {
                    printf("Errore, il numero di posti deve essere positivo.\n");
                    break;
                }
               
                rgstr = new_book(n, rgstr, tmp_n);
                break;
            case '+': // + k name --> book(k, name)
                scanf(" %d ", &k);
                
                if (k < 0) {
                    printf("Errore, la posizione deve essere positiva.\n");
                    break;
                }

                str = malloc(10 * sizeof(char));

                while((c = getchar()) != '\n') {
                    if (!isalpha(c)) {
                        break;
                    }
                    if (count >= size) {
                        size += 10;
                        str = realloc(str, size);
                    }
                    str[count++] = c;
                }
                str[count] = '\0';
                book(k, n, str, rgstr);
                size = 10;
                count = 0;
                free(str);
                break;
            case '-': // - k --> cancel(k)
                scanf(" %d", &k);
                cancel(k, n, rgstr);
                break;
            case 'm': // m from to ---> move from to
                scanf(" %d", &from);
                scanf(" %d", &to);
                move(rgstr, from, to, n);
                break;
            case 'p': // p ---> printBook()
                print_book(rgstr, n);
                break;
        }
    }
}

/* This function creates a new book, and if it altready exists delete it
 *
 * Pre-condition: n > 0 .
 * Side effects: *tmp_n will be modified, standard output can be modified.
 * Post-condition: This funciton return address of register.
 */

char **new_book(int n, char **r, int *tmp_n) {
    if(r != NULL) {
        for(int i = 0; i < *tmp_n; i++) {
            free(r[i]);
        }
        free(r);
    }

    *tmp_n = n;
    r = malloc(n * sizeof(char *));
    
    if(r == NULL) {
        printf("Errore durante l'allocazione!\n");
        exit(EXIT_FAILURE);
    }

    for(int i = 0; i < n; i++) {
        *(r + i) = malloc(1 * sizeof(char)); // *(r + i) = r[i]
        if(r == NULL) {
            printf("Errore durante l'allocazione!\n");
            exit(EXIT_FAILURE);
        }
    }

    for(int j = 0; j < n; j++) {
        r[j][0] = '\0';
    }

    return r;
}


/* This function makes a booking
 * Side effects: standard output can be modified. 
 * Pre-condition: k > 0, name != ' '.
 */

void book(int k, int n, char *name, char **r) {
    if(k >= n || (r[k][0] != '\0')) {
        printf("Errore, posto inesistente o gia' occupato!\n");
    } else {
        r[k] = realloc(r[k], strlen(name));
        if(r == NULL) {
            printf("Errore durante l'allocazione!\n");
            exit(EXIT_FAILURE);
        }
        strcpy(r[k], name);
    }
}


/* This function clears a place of book
 *
 * Pre-condition: k > 0 && k < n.
 * Side effects: r[k] will be contain '\0', standard output can be modified.
 * Post-condition: place k become free 
 */

void cancel(int k, int n, char **r) {
     if(k >= n || r[k][0] == '\0') {
        printf("Errore, posto inesistente o gia' libero!\n");
    } else {
        free(r[k]);
        r[k] = malloc(1 * sizeof(char));
        if(r == NULL) {
            printf("Errore durante l'allocazione!\n");
            exit(EXIT_FAILURE);
        }
        r[k][0] = '\0';
    }   
}


/* This function moves the content of 'from' position to 'to' position in the book
 *
 * Pre-condition: from > 0 && from < n && from != to && to > 0 && to < n && to != from..
 * Side effects: standard output can be modified.
 * Post-condition: r[from] moved to r[to] 
 */

void move(char **r, int from, int to, int n) {
    if (from >= n || to >= n) {
        printf("Errore, posizioni indicate sbagliate!\n");
    } else if (r[from][0] == '\0') {
        printf("La posizione %d e' libera, non e' possibile spostarla in %d\n", from, to);
    
    } else if(r[to][0] != '\0') {
        printf("La posizione %d e' gia' occupata, non e' possibile effettuare lo spostamento.\n", to);
    } else {
        book(to, n, r[from], r);
        cancel(from, n, r);
    }
}


/* This function prints the status of bool
 *
 * Side effects: standard output can be modified.
 */

void print_book(char **r, int n) {
    printf("REGISTER [%d..%d]\n", (n-n), (n-1));
    for(int i = 0; i < n; i++) {
        printf("%d --> %s\n", i, r[i]);
    }
}
