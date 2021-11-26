#include "stack.h"
#include <stdio.h>
#include <stdlib.h>
#define DIM 100
/* Usato per stampare utilizzando i generici
#define print_any(X) _Generic((X), int: print_int, \
                                   char: print_char, \
                                   float: print_float, \
                                   default: print_unknown)(X)
int print_int(int i) {
    return printf("%d", i);
}

int print_float(float f) {
    return printf("%f", f);
}

int print_char(char c) {
    return printf("%c", c);
}

int print_unknown() {
    return printf("ERROR: unknown type\n");
}

*/
Item stack[DIM];
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

Item top(void) {
    if(is_empty()) { 
        printf("Errore, la pila e' vuota e non e' possibile estrarre alcun valore\n");
        exit(EXIT_FAILURE);
    }
    return stack[head];
}

Item pop(void) {
    if(is_empty()) {
       printf("Errore, la pila e' vuota e non e' possibile estrarre alcun valore\n");
       exit(EXIT_FAILURE);
    }
    Item val = top();
    head--;
    return val;
}

void push(Item n) {
    if(is_full()) {
        printf("Errore, la pila e' piena\n");
        return;
    }
	head++;
	stack[head] = n;
}

void print_stack(void) {
    for(int i = head; i >= 0; i--) {
		printf("chiave %d --- ", i);
        print_item(stack[i]); //print_any(stack[i]);
        printf("\n");
	}
}

