#include "list.h"
#include <stdio.h>
#include <stdlib.h>

typedef struct node *Node;
/* Stampa usando i generici
#define print_any(X) _Generic((X), int: print_int, \
                                   char: print_char, \
                                   float: print_float, \
                                   default: print_unknown)(X)
int print_int(int i) {
    return printf("%d ", i);
}

int print_float(float f) {
    return printf("%f ", f);
}

int print_char(char c) {
    return printf("%c ", c);
}

int print_unknown() {
    return printf("ERROR: unknown type\n");
}
*/

void list_insert_void(Item n, Node *l) { // e' come fare struct node **l
    Node new_node = malloc(sizeof(Node));
    new_node -> info = n;
    new_node -> next = *l;

}

Node list_insert(Item n, Node l) {
    Node new_node = malloc(sizeof(Node));
    new_node -> info = n;
    new_node -> next = l;
    return new_node; 
}

Node list_search(Item n, Node l) {
    while(l != NULL && l -> info != n) {
        l = l -> next;
    }

    return l;
}

Node list_search_rec(Item n, Node l) {
    if (l -> info == n) {
        return l;
    }

    if (l -> next == NULL) {
        return NULL;
    }

    l = l -> next;
    list_search_rec(n, l);
    return NULL;
}

Node list_delete(Item n, Node l) {
    Node curr, prev;
    for ( curr = l, prev = NULL; curr != NULL; prev = curr, curr = curr -> next ) {
        if (curr -> info == n )
            break;
    }
    if (curr == NULL)
        return l;
    if (prev == NULL) {
        l = l -> next; 
    }
    else
        prev -> next = curr -> next;
    free(curr);
    return l;
}

void print_list(Node l) {
    for(;l != NULL;) {
        print_item(l -> info); //print_any(l -> info);
        l = l -> next;
    }
    printf("\n");
}

void print_list_inv(Node l) {
    if(l -> next != NULL) {
        print_list_inv(l -> next); 
    }
    print_item(l -> info); //print_any(l -> info);
}

Item *list_to_array(Node l, int *dim) {
    int i = 0, size = 10;
    Item *arr = malloc(size * sizeof(int));

    for(i = 0; l != NULL; i++) {
        if (i >= size) {
            i += 5;
            arr = realloc(arr, i);
        }
        arr[i] = l -> info;
        l = l -> next;
    }
    *dim = i;
    return arr;
}

void list_destroy(Node *l) {
    Node current = *l;
    Node next = NULL;
    while(current != NULL) {
        next = current -> next;
        free(current);
        current = next;
    }
    *l = current;
}

void list_destroy_rec(Node *l) {
    Node current = *l;
    if(current != NULL) {
        list_destroy_rec(&(current -> next));
    }
    free(current);
    *l = NULL;
}

Node o_list_insert(Item n, Node l) {
    Node curr = l;
    Node prec = NULL;
    for(; curr != NULL; prec = curr, curr = curr -> next) { 
        if(curr -> info > n) {
            Node ins = malloc(sizeof(Node));
            ins -> info = n;
            ins -> next = curr;
            prec -> next = ins;
            break;
        }
    }
    return l;
}

Node o_list_search(Item n, Node l) {
    while(l != NULL) {
        if(l -> info == n) {
            return l;
        } else if(l -> info > n) {
            return NULL;
        }
        l = l -> next;
    }
    return l;
}

