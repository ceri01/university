#ifndef ALBERI_BINARI_H
#define ALBERI_BINARI_H
#include "../../2021-11-12/item/item.h"

struct bit_node {
    Item item;
    struct bit_node *l, *r;
};

typedef struct bit_node *Bit_node;

Bit_node bit_new(Item i);

void bit_destroy(Bit_node t);

Bit_node bit_right(Bit_node t);

Bit_node bit_left(Bit_node t);

Item bit_item(Bit_node t);

void bit_preorder(Bit_node t);

void bit_inorder(Bit_node t);

void bit_postorder(Bit_node t);

void bit_print_node(Bit_node t);

void bit_print_as_summary(Bit_node p, int spaces);

Bit_node bit_arr_to_tree(Item a[], int size, int i);

#endif

