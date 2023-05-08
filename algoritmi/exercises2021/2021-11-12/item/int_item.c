#include "item.h"
#include <stdio.h>
#include <stdlib.h>

Item atoitem(char *str) {
    return atoi(str);
}

void print_item(Item i) {
    printf("%d", i);
}


