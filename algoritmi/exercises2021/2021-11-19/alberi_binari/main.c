#include "../../2021-11-12/item/item.h"
#include "alberi_binari.h"
#include <stdio.h>

int main() {
    Item bt[10] = {4, 1, 9, 3, 8, 11, 23, 56, 44, 21};
   	Bit_node tree = bit_arr_to_tree(bt, 10, 0);	
    bit_print_as_summary(tree, 0);
}
