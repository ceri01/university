#include "list.h"
#include <stdio.h>
#include <stdlib.h>

typedef struct node *Node;

void list_insert_void(int n, Node *l) { // e' come fare struct node **l
    Node new_node = malloc(sizeof(Node));
    new_node -> info = n;
    new_node -> next = *l;

}

Node list_insert(int n, Node l) {
    Node new_node = malloc(sizeof(Node));
    new_node -> info = n;
    new_node -> next = l;
    return new_node; 
}

Node list_search(int n, Node l) {
    while(l != NULL && l -> info != n) {
        l = l -> next;
    }

    return l;
}

Node list_search_rec(int n, Node l) {
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

Node list_delete(int n, Node l) {
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
        printf("%d " , l -> info);
        l = l -> next;
    }
    printf("\n");
}

void print_list_inv(Node l) {
    if(l -> next != NULL) {
        print_list_inv(l -> next); 
    }
    printf("%d ", l -> info);
}

int *list_to_array(Node l, int *dim) {
    int i = 0, size = 10;
    int *arr = malloc(size * sizeof(int));

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

Node o_list_insert(int n, Node l) {
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

Node o_list_search(int n, Node l) {
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
