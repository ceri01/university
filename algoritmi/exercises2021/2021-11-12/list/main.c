#include "list.h"
#include <stdio.h>
#include <stdlib.h>
/* Scan con generici non funzionante
#define scan_any(X) _Generic((X), int: scan_int, \
                                  float: scan_float, \
                                  char: scan_char)(X)

int scan_int(int i) {
    scanf("%d", &i);
    return i;
}
int scan_float(float f) {
    printf("%f", f);
    scanf("%f", &f);
    return f;
}
int scan_char(char c) {
    scanf("%c", &c);
    return c;
}
*/
typedef struct node *Node;

typedef struct {
    Node head;
    int count;
} Set;

int main() {
    Set *list = malloc(sizeof(Set));
    list -> head = NULL;
    list -> count = 0;
    Node index_list =  list -> head;
    char ch;
    Item n;

    while((ch = getchar()) != 'f') {
        switch (ch) {
            case '+': 
                n = scan_any(n);
                index_list = list_search(n, list -> head);
                if(index_list == NULL) {
                    list -> head = list_insert(n, list -> head);
                    list -> count++;
                }
                break;
            case '-':
                n = scan_any(n);
                index_list = list_search(n, list -> head);
                if(index_list == NULL) {
                    break;
                } else {
                    list -> head = list_delete(n, list -> head);
                    list -> count--;
                }
                break;
            case '?': 
                n = scan_any(n);
                index_list = list_search(n, list -> head);
                if(index_list == NULL) {
                    printf("%d non appartiene all'insieme\n", n);
                } else {
                    printf("%d appartiene all'insieme\n", n);
                }
                break;
            case 'c':
                printf("Numero elementi dell'insieme = %d\n", list -> count);
                break;
            case 'p':
                print_list(list -> head);
                break;
            case 'd':
                list_destroy(&list -> head);
                list -> count = 0;
                break;
        }
    }
}
