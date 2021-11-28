#include "b_list.h"
#include <stdio.h>
#include <stdlib.h>
#include "../../2021-11-12/item/item.h"

typedef struct node *Node;

typedef struct {
    Node head;
    int count;
} Set;

int main() {
    Set *list = malloc(sizeof(Set));
    list -> head = NULL;

    list -> count = 0;	
	list -> head = list_delete(7, list -> head);
	print_list(list -> head);
}
