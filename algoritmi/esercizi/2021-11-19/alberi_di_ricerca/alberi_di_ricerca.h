#ifndef ALBERI_BINARI_DI_RICERCA_H
#define ALBERI_BINARI_DI_RICERCA_H
#include "../../2021-11-12/item/item.h"

struct occurrence {
	char* word; // key
	int n; // val
};

typedef struct occurrence Occurrence;

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

void bist_order_print(Bit_node p);

void bist_inv_order_print(Bit_node p);

Item bist_search(Bit_node r, char* k);

void bist_insert(Bit_node *q, Item item);

int bist_delete(Bit_node *r, char* k);

#endif

