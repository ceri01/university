#include "../../2021-11-12/item/item.h"
#include "alberi_binari.h"
#include <stdio.h>
#include <stdlib.h>

Bit_node bit_new(Item i) {
    Bit_node n = malloc(sizeof(struct bit_node));
    n -> item = i;
    n -> l = NULL;
    n -> r = NULL;
    return n;
}

void bit_destroy(Bit_node t) {
    free(t);
}

Bit_node bit_right(Bit_node t) {
   return t -> r;
}

Bit_node bit_left(Bit_node t) {
   return t -> l;
}

Item bit_item(Bit_node t) {
   return t -> item;
}

void bit_preorder(Bit_node t) {
    if(t != NULL) {
        bit_print_node(t);
        bit_preorder(t -> l);
        bit_preorder(t -> r);

    }
}

void bit_inorder(Bit_node t) {
    if(t != NULL) {
        bit_inorder(t -> l);
        bit_print_node(t);
        bit_inorder(t -> r);
    }
} 

void bit_postorder(Bit_node t) {
    if(t != NULL) {
        bit_postorder(t -> l);
        bit_postorder(t -> r); 
        bit_print_node(t);
    } 
}

void bit_print_node(Bit_node t) {
    print_item(t -> item);
    printf(" ");
}


void bit_print_as_summary(Bit_node p, int spaces) { 
    for(int i = 0; i < spaces; i++) {
        printf(" ");
    }
    if(p != NULL) {
        printf("* ");
        bit_print_node(p); 
        printf("\n");
        if(p -> l == NULL && p -> r == NULL) {
            return;            
        }
        bit_print_as_summary(p -> l, spaces + 4);
        bit_print_as_summary(p -> r, spaces + 4);
    } else {
        printf("*\n");
    }
}

