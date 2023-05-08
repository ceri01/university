
#ifndef LIST_H
#define LIST_H

struct node {
    int info;
    struct node *next;
};

/* Permette di inserire un elemento n nella lista */
void list_insert_void(int n, struct node **l);

/* Permette di inserire un elemento n nella lista, ritorna il puntatore al nuovo elemento inserito (ovvero la testa della lista) */
struct node *list_insert(int n, struct node *l);

/* Permette di cercare un elemento n nella lista, ritorna il puntatore
 * al nodo cercato se presente, NULL altrimenti */
struct node *list_search(int n, struct node *l);

/* Permette di cercare un elemento n nella lista in modo ricorsivo, ritorna il puntatore al nodo cercato se presente, NULL altrimenti */
struct node *list_search_rec(int n, struct node *l);

/* Elimina il nodo specificato se presente. abbiamo 3 casi: se non ci
 * sono elementi da eliminare viene ritornato lo stesso l preso in ingresso,
 * se invece viene eliminato il primo elemento viene posto il secondo elemento
 * della lista come testa e poi ritornato. infine se viene eliminato un altro elemento
 * in mezzo alla lista verra' ritornata la testa */
struct node *list_delete(int n, struct node *l);

/* Permette di stampare la lista passata */
void print_list(struct node *l);

/* Permette di stampare gli elenenti della lista al contrario */
void print_list_inv(struct node *l);

/* Restituisce il puntatore ad un array allocato dinamicamente contenente gli elementi della lista, in piu' il puntatore ad int passato prendera il valore della dimensione dell array */
int *list_to_array(struct node *l, int *dim);

/* Effettua la cancellazione della lista liberando lo spazio */
void list_destroy(struct node **l);

/* Effettua la cancellazione della lista liberando lo spazio (in modo ricorsivo) */
void list_destroy_rec(struct node **l);

/* Inserisce il valore n nella lista mantenendo l'ordine, viene restituito l'indirizzo della testa */
struct node *o_list_insert(int n, struct node *l);

/* Cerca il valore passato n nella lista ordinata, viene restituito l'indirizzo della testa */
struct node *o_list_search(int n, struct node *l);

#endif
