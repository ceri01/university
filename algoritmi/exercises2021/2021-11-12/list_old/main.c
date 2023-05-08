#include "list.h"
#include <stdio.h>
#include <stdlib.h>

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
    int n;

    while((ch = getchar()) != 'f') {
        switch (ch) {
            case '+':
                scanf(" %d", &n);
                index_list = list_search(n, list -> head);
                if(index_list == NULL) {
                    list -> head = list_insert(n, list -> head);
                    list -> count++;
                }
                break;
            case '-':
                scanf(" %d", &n);
                index_list = list_search(n, list -> head);
                if(index_list == NULL) {
                    break;
                } else {
                    list -> head = list_delete(n, list -> head);
                    list -> count--;
                }
                break;
            case '?':
                scanf(" %d", &n);
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
