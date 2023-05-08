#include "../../2021-11-12/item/item.h"
#include "alberi_di_ricerca.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

Bit_node bit_new(Item i) {
    Bit_node n = malloc(sizeof(struct bit_node));
    n -> item = i;
    n -> l = NULL;
    n -> r = NULL;
    return n;
}

void bit_destroy(Bit_node t) {
    free(t);
}

Bit_node bit_right(Bit_node t) {
   return t -> r;
}

Bit_node bit_left(Bit_node t) {
   return t -> l;
}

Item bit_item(Bit_node t) {
   return t -> item;
}

void bit_preorder(Bit_node t) {
    if(t != NULL) {
        bit_print_node(t); 
		printf("\n");
		bit_preorder(t -> l);
        bit_preorder(t -> r);

    }
}

void bit_inorder(Bit_node t) {
    if(t != NULL) {
        bit_inorder(t -> l);
        bit_print_node(t);
		printf("\n");
        bit_inorder(t -> r);
    }
} 

void bit_postorder(Bit_node t) {
    if(t != NULL) {
        bit_postorder(t -> l);
        bit_postorder(t -> r); 
        bit_print_node(t);
		printf("\n");
    } 
}

void bit_print_node(Bit_node t) {
    print_item(t -> item);
    printf(" ");
}


void bit_print_as_summary(Bit_node p, int spaces) { 
    for(int i = 0; i < spaces; i++) {
        printf(" ");
    }
    if(p != NULL) {
        printf("* ");
        bit_print_node(p); 
        printf("\n");
        if(p -> l == NULL && p -> r == NULL) {
            return;            
        }
        bit_print_as_summary(p -> l, spaces + 4);
        bit_print_as_summary(p -> r, spaces + 4);
    } else {
        printf("*\n");
    }
}

void bist_insert(Bit_node *r, Item item) {
	if(*r == NULL) {
		Bit_node new_node = malloc(sizeof(struct bit_node));
		new_node -> item = item;
		new_node -> l = NULL;
		new_node -> r = NULL;
		*r = new_node;
		return;
	} else {
		if(strcmp((*r) -> item.word, item.word) > 0) {
			bist_insert(&((*r) -> l), item);  
		}
		if(strcmp((*r) -> item.word, item.word) < 0) {
			bist_insert(&((*r) -> r), item);  
		}
	}
}


Item bist_search(Bit_node r, char* k) {
	while(r != NULL) {
		if(strcmp(r -> item.word, k) > 0) {
			r = r -> l;
		} else if(strcmp(r -> item.word, k) < 0) {
			r = r -> r;  
		} else {
			return r -> item;
		}
	}
	Item i = {NULL, -1};
	printf("Oggetto non trovato\n");
	return i;
}

void bist_delete(Bit_node *r, char* k) {
	Bit_node *curr_node = r;
    Bit_node father = NULL;
    
    while (*curr_node != NULL && strcmp((*curr_node) -> item.word, k) != 0) {
        father = *curr_node;
        if(strcmp(k, (*curr_node) -> item.word) < 0) {
            *curr_node = (*curr_node) -> l;
        } else {
            *curr_node = (*curr_node) -> r;
        }
    }
    if(*curr_node != NULL) {
        if((*curr_node) -> l == NULL) {
            if(father != NULL) {
                if(strcmp((*curr_node) -> item.word, father -> item.word) < 0) {
                    father -> l = (*curr_node) -> r;
                } else {
                    father -> r = (*curr_node) -> r;
                }
            } else {
                *r = (*r) -> r;
            }
        } else if((*curr_node) -> r == NULL) {
            if(father != NULL) {
                if(strcmp((*curr_node) -> item.word, father -> item.word) < 0) {
                    father -> l = (*curr_node) -> l;
                } else {
                    father -> r = (*curr_node) -> l;
                }
            } else {
                *r = (*r) -> l;
            }
        } else {
            Bit_node t = *curr_node;
            Bit_node m = (*curr_node) -> l;
            while(m -> r != NULL) {
                t = m;
                m = m -> r;
            }
            (*curr_node) -> item = m -> item;
            if(t == *curr_node) {
                (*curr_node) -> l = m -> l;
            } else {
                t -> r = m -> l;
            }
        
        }
	}
}


void bist_order_print(Bit_node p) {
	if(p != NULL) {
		bit_inorder(p);	
	}
}
/*
void bist_inv_order_print(Bit_node p) {
	
}
*/





