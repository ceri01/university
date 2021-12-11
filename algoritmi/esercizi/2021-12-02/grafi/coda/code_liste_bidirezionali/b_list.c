#include "b_list.h"
#include <stdio.h>
#include <stdlib.h>

Node list_insert(int n, Node l) { // e' come fare struct node **l
    Node new_node = malloc(sizeof(Node));

	new_node -> info = n;
    new_node -> next = l;
	new_node -> prev = NULL;
	if(l != NULL) {
		l -> prev = new_node;
	}
	return new_node;
}

Node list_search(int n, Node l) {
    while(l != NULL && l -> info != n) {
        l = l -> next;
    }
    return l;
}

Node list_delete(int n, Node l) {
    Node new_head = l;
	for (; l != NULL; l = l -> next) { 
        if (l -> info == n ) { 
			break;
		}
    }
    if (l == NULL)
        printf("Elemento non presente nella lista\n");
		return l;
    if (l -> prev == NULL) {
		new_head = l -> next;
		free(l);
		return new_head;
    } else {
		l -> prev -> next = l -> next;
	}
    free(l);
    return new_head;
}

void print_list(Node l) {
    for(;l != NULL;) {
        print_item(l -> info); //print_any(l -> info);
        printf(" ");
		l = l -> next;
    }
    printf("\n");
}

void list_destroy(Node *l) {
    Node current = *l;
    Node next = NULL;
    while(current != NULL) {
        next = current -> next;
        free(current);
        current = next;
    }
    *l = current;
}




