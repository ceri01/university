#ifndef LIST_H
#define LIST_H
#include "../../2021-11-12/item/item.h"

struct node {
    Item info;
    struct node *next;
	struct node *prev;
};

typedef struct node *Node;

/* Permette di inserire un elemento n nella lista */
Node list_insert(Item n, Node l);

/* Permette di cercare un elemento n nella lista, ritorna il puntatore
 * al nodo cercato se presente, NULL altrimenti */
Node list_search(Item n, Node l);

/* Elimina il nodo specificato se presente. abbiamo 3 casi: se non ci
 * sono elementi da eliminare viene ritornato lo stesso l preso in ingresso,
 * se invece viene eliminato il primo elemento viene posto il secondo elemento
 * della lista come testa e poi ritornato. infine se viene eliminato un altro elemento
 * in mezzo alla lista verra' ritornata la testa */
Node list_delete(Item n, Node l);

/* Permette di stampare la lista passata */
void print_list(Node l);

/* Effettua la cancellazione della lista liberando lo spazio */
void list_destroy(Node *l);

#endif
