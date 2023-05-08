#include "../../2021-11-12/item/item.h"
#include "alberi_di_ricerca.h"
#include <stdio.h>

int main() {
   	Bit_node root = NULL;
	Item i1;
	i1.word = "basato";
	i1.n = 2;
	Item i2;
	i2.word = "casa";
	i2.n = 3;
	Item i3;
	i3.word = "aratro";
	i3.n = 1;	
    bist_insert(&root, i1);
    bist_insert(&root, i2);
    bist_insert(&root, i3);
	bist_order_print(root);
	printf("\n");
	Item kek = bist_search(root, "casa");
	printf("%s --> %d\n", kek.word, kek.n);
    bist_delete(&root, "basato");
    bist_order_print(root);
}
