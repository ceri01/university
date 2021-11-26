#include "stack.h"
#include <stdio.h>
#include <stdlib.h>
#define DIM 10

int stack[DIM];
int head = -1;

void make_empty(void) {
    head = -1;
}

int is_empty(void) {
    if(head == -1) {
        return 1;
    }
    return 0;
}

int is_full(void) {
    if(head >= DIM - 1) {
		return 1;
	}
	return 0;
}

int top(void) {
    if(is_empty()) { 
        printf("Errore, la pila e' vuota e non e' possibile estrarre alcun valore\n");
        exit(EXIT_FAILURE);
    }
    return stack[head];
}

int pop(void) {
    if(is_empty()) {
       printf("Errore, la pila e' vuota e non e' possibile estrarre alcun valore\n");
       exit(EXIT_FAILURE);
    }
    int val = top();
    head--;
    return val;
}

void push(int n) {
    if(is_full()) {
        printf("Errore, la pila e' piena\n");
        return;
    }
	head++;
	stack[head] = n;
}

void print_stack(void) {
    for(int i = head; i >= 0; i--) {
		printf("chiave %d --- %d\n", i, stack[i]);
	}
}

