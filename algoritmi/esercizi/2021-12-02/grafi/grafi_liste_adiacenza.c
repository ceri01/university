#include "grafi.h"
#include <stdio.h>
#include <stdlib.h>
#include "coda/code_liste_bidirezionali/l_queue.h"
#include <time.h>

struct list_node {
	struct list_node *next;
	int v;
};

struct graph {
	int n, m;
	struct list_node **A; 
};

Graph graph_new(int n) {
	Graph g = malloc(sizeof(struct graph));
	g -> n = n;
	g -> A = calloc(n, sizeof(struct list_node*));
	for(int i = 0; i < n; i++) {
		g -> A[i] = NULL;
	}
	return g;
}

void graph_destroy(Graph g) {
	if(g == NULL) {
        printf("Grafo nullo.\n");
        return;
    }
    for(int i = 0; i < g -> n; i++) {
		struct list_node *tmp, *curr;
		curr = g -> A[i];
        if(curr == 0) {
            continue;
        }
		while(curr -> next != NULL) {
			tmp = curr -> next;
			free(curr);
			curr = tmp;
		}
	}
	free(g -> A);
	free(g);
}

void graph_edge_insert(Graph g, int v, int w) {
    if(g -> n <= 1 || g == NULL) {
        printf("Grafo non esistente oppure formato da un solo nodo\n");
        return;
    }
    if(v == w) {
        printf("Impossibile aggiungere un arco da un nodo a se stesso.\n");
        return;
    }
    struct list_node *ptr1, *ptr2;
    struct list_node *prev = g -> A[v];
    struct list_node *el1 = malloc(sizeof(struct list_node));
    el1 -> v = w;
    el1 -> next = NULL;
    struct list_node *el2 = malloc(sizeof(struct list_node));	 
    el2 -> v = v;
    el2 -> next = NULL;
    
    ptr1 = g -> A[v];
    ptr2 = g -> A[w];
    
    while(ptr1 != NULL) {
        if(ptr1 -> v == w) {
            printf("Arco gia' presente.\n");
            return;
        }
        prev = ptr1;
        ptr1 = ptr1 -> next;
    }
    if(prev == NULL) {
        g -> A[v] = el1;    
    } else {
        prev -> next = el1;
    }

    prev = g -> A[w];

    while(ptr2 != NULL) {
        prev = ptr2;
        ptr2 = ptr2 -> next;
    }
    if(prev == NULL) { 
        g -> A[w] = el2;    
    } else {
        prev -> next = el2;
    }
    return;
}

Graph graph_read() {
    int n, np, na;
    printf("Inserisci il numero di nodi del grafo.\n");
    scanf("%d", &n);
    if(n <= 0) {
        printf("Errore, impossibile creare un grafo con meno di 1 nodo.\n");
    }
    Graph g = graph_new(n);
    printf("Inserisci un arco (nodi da 0 a %d).\n", (n - 1));
    scanf(" %d %d", &np, &na); 
    while(np > 0 || na > 0) {
        graph_edge_insert(g, np, na); 
        printf("Inserisci un arco (nodi da 0 a %d).\n", (n - 1));
        scanf(" %d %d", &np, &na);
    }
    return g;
}

void graph_print(Graph g) {
    if(g == NULL) {
        printf("Il grafo non esiste, ha valore NULL\n");
        return;
    }
	for(int i = 0; i < g -> n && g -> A[i] != NULL; i++) {
		struct list_node *tmp = g -> A[i];
		printf("[%d] ", i);
		printf("%d", g -> A[i] -> v);
		while(tmp -> next != NULL) {
			tmp = tmp -> next;
			printf(" -> %d", tmp -> v);
		}
		printf("\n");
	}
}

void dfs(Graph g) {
    int *vis = calloc(g -> n, sizeof(int));
    srand(time(NULL));
    for(int i = (rand() % g -> n); i < g -> n; i++) {
        if(vis[i] != 1) {
            printf("%d ", i);
            dfs_rec(g, i, vis);
        }
    }
}

void dfs_rec(Graph g, int i, int *vis) {
    struct list_node *tmp = g -> A[i];
    vis[i] = 1;
    while(tmp != NULL) {
        if(vis[tmp -> v] != 1) {
            printf("%d ", tmp -> v);
            dfs_rec(g, tmp -> v, vis);
        }
        tmp = tmp -> next;
    }
}

void bfs(Graph g) {
}
