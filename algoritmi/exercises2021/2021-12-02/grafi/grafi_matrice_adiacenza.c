#include "grafi.h"
#include <stdio.h>
#include <stdlib.h>
#include "coda/code_liste_bidirezionali/l_queue.h"

struct graph {
    int n, m; /* n = num vertici, m = num lati */
    int **A; /* matrice di adiacenza */
};

Graph graph_new(int n) {
    if(n < 0) {
        printf("Errore, il numero di nodi inserito e' invalido.\n");
        exit(EXIT_FAILURE);
    }
    Graph gr = malloc(sizeof(struct graph));
    if(gr == NULL) {
        printf("Errore, errore nella malloc.\n");
        exit(EXIT_FAILURE);
    }

    gr -> A = calloc(n, sizeof(int*));
    if(gr -> A == NULL) {
        printf("Errore, errore nella malloc.\n");
        exit(EXIT_FAILURE);
    }

    for(int i = 0; i < n; i++) {
        gr -> A[i] = calloc(n, sizeof(int));
        if(gr -> A[i] == NULL) {
            printf("Errore, errore nella malloc.\n");
            exit(EXIT_FAILURE);
        }
        for(int k = 0; k < n; k++) {
            gr -> A[i][k] = 0;
        }
    }
    gr -> n = n;
    return gr;
}

void graph_destroy(Graph g) {
    if(g == NULL) {
        printf("La matrice di adiacenza non e' stata ancora creata.\n");
        return;
    }
    for(int i = 0; i < g -> n; i++) {
        free(g -> A[i]);
    }
    free(g -> A);
    free(g);
}

void graph_edge_insert(Graph g, int v, int w) {
    if(g -> m >= (((g -> n) * (g -> n)) - (g -> n))/2) {
        printf("Non e' possibile inserire ulteriori archi.\n");
        return;
    }
    if((v >= g -> n || v < 0) || (w >= g -> n || w < 0)) {
        printf("Impossibile creare l'arco, uno dei due nodi inseriti non esiste.\n");
        return;
    } 
    g -> A[v][w] = 1;
    g -> A[w][v] = 1;
    (g -> m)++;
    return;
}

Graph graph_read() {
    int n1, n2, limit;
    Graph g;
    printf("Inserire numero nodi e numero archi da inserire\n"); 
    // inizio: n1 numero nodi, n2 numero archi da inserire
    scanf(" %d %d", &n1, &n2);
    if(n1 > 0 && n2 > ((n1 * n1) - n1)/2) {
        printf("Numero archi superiore al massimo, numero di archi settat = numero max di archi.\n");
        n2 = ((n1 * n1) - n1)/2;
    }
    g = graph_new(n1);
    g -> m = 0;
    limit = n2;

    // inserimento archi
    while(g -> m < limit) {
        printf("Inserisci un nuovo arco.\n");
        scanf(" %d %d", &n1, &n2);
        graph_edge_insert(g, n1, n2);
        if(g -> m >= (((g -> n) * (g -> n)) - (g -> n))/2) {
            return g;
        }
    }
    return g;
}

void graph_print(Graph g) {
    for(int i = 0; i < g -> n; i++) {
        for(int k = i; k < g -> n; k++) {
            if(g -> A[i][k] == 1) {
                printf("%d <---> %d\n", i, k);
            }
        }
    }
}

void bfs(Graph g) {
    Queue q = NULL;
    q = create_queue(q);
    int *vis = calloc(g -> n, sizeof(int));
    int i = 0;

    vis[0] = 1;
    enqueue(q, 0);

    while(!is_empty(q)) {
        i = dequeue(q);
        for (int j = 0; j < g -> n; j++) {
            if(vis[j] == 0 && g -> A[i][j] == 1) {
                vis[j] = 1;
                enqueue(q, j);
            }
        }
        printf("%d ", i);
    }
    printf("\n");
}

void dfs(Graph g) {
    int *vis = calloc(g -> n, sizeof(int));
    for(int i = 0; i < g -> n; i++) {
        if(vis[i] != 1) {
            vis[i] = 1;
            printf("%d ", i);
            dfs_rec(g, i, vis);
        }
    }
    printf("\n");
}

void dfs_rec(Graph g, int i, int *vis) {
    for(int j = 0; j < g -> n; j++) {
        if(g -> A[i][j] == 1 && vis[j] != 1) {
            vis[j] = 1;
            printf("%d ", j);
            dfs_rec(g, j, vis);
        }
    }
}
