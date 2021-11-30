#include "../../2021-11-12/item/item.h"
#include "alberi_binari.h"
#include <stdio.h>

int main() {
    Bit_node bt;
    bt = bit_new(5);
    bt -> r = bit_new(6);
    bt -> l = bit_new(1); 
    bt -> r -> l = bit_new(2);
    bt -> r -> r = bit_new(7);
    bt -> l -> l = bit_new(4);
    //bt -> l -> r = bit_new(9);
    bit_print_as_summary(bt, 0);
}
