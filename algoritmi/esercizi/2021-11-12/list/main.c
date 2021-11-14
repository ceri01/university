#include "list.h"
#include <stdio.h>

int main() {
    struct node *first = NULL;
    struct node *sr = NULL;
    first = list_insert(11, first);
    first = list_insert(8, first);
    first = list_insert(7, first);
    first = list_insert(6, first);
    first = list_insert(2, first);
    first = list_insert(1, first);
    
    first = o_list_insert(4, first);
    sr = o_list_search(11, first);

    print_list(sr);
}
