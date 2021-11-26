#include <stdio.h>
#include <stdlib.h>
#include "../list/list.h"

typedef struct node *Node;

typedef struct {
    Node head;
    Node last;
    int dim;
} Stack;

Stack stk = {NULL, NULL, 10};

void make_empty(void) {
    if(stk.head != NULL) {
        list_destroy(&stk.head);
        stk.last = NULL;
    } else {
        printf("La pila e' gia' vuota.\n");
    }
}

int is_empty(void) {
    if(stk.head == NULL) {
        return 1;
    }
    return 0;
}

int is_full(void) {
    int count = 0;
    Node tmp = stk.head;
    if(is_empty()) {
        return 0;
    }
    while(tmp != NULL) {
        count++;
        tmp = tmp -> next;
    }
   
    if (count < stk.dim) {
        return 0;
    }
    return 1;
}

Item top(void) {
    if(is_empty()) { 
        printf("Errore, la pila e' vuota e non e' possibile estrarre alcun valore\n");
        exit(EXIT_FAILURE);
    }
    return stk.head -> info;
}

Item pop(void) {
    if(is_empty()) {
       printf("Errore, la pila e' vuota e non e' possibile estrarre alcun valore\n");
       exit(EXIT_FAILURE);
    }
    Item val = top();
    stk.head = list_delete(val, stk.head);
    return val;
}

void push(Item n) {
    if(is_full()) {
        printf("Errore, la pila e' piena\n");
        return;
    }
    stk.head = list_insert(n, stk.head);
}

void print_stack(void) {
    Node tmp = stk.head;
    print_list(tmp);
}

