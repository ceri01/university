#include "item.h"
#include <stdio.h>
#include <stdlib.h>

Item atoitem(char *str) {
    return atof(str);
}

void print_item(Item i) {
    printf("%f", i);
}


