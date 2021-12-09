#include "grafi.h"
#include <stdio.h>
#include <stdlib.h>

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
	for(int i = 0; i < g -> n; i++) {
		struct list_node tmp, curr;
		curr = g -> A[i];
		while(g -> A[i] -> next != NULL) {
			tmp = curr -> next;
			free(curr);
			curr -> tmp;
		}
	}
	free(g -> A);
	free(g);
}

void graph_edge_insert(Graph g, int v, int w) {
	if(v == w) {
		printf("Non e' possibile collegare un nodo a se stesso.\n");
		break;
	}
	if(g -> A[v] == NULL) {
		g -> A[v] = malloc(sizeof(list_node));
		g -> A[v] -> v = w;
		g -> A[v] -> next = NULL;
		if(g -> A[w] == NULL) {
			g -> A[w] = malloc(sizeof(list_node));
			g -> A[w] -> v = v;
			g -> A[w] -> next = NULL;
		} else {	
			struct list_node tmp;
			while(g -> A[w] -> next != NULL) {
				tmp = g -> A[w] -> next;
			}
			tmp = malloc(sizeof(list_node));
			tmp -> v = v;
			tmp -> next = NULL;
		}
		break;
	} else {
		struct list_node tmp;
		while(g -> A[v] -> next != NULL) {
			if(g -> A[v] -> v == w) {
				printf("Collegamento gia' presente.\n");
			}
			tmp = g -> A[v] -> next;
		}
		tmp = malloc(sizeof(list_node));
		tmp -> v = w;
		tmp -> next = NULL;
		if(g -> A[w] == NULL) {
			g -> A[w] = malloc(sizeof(list_node));
			g -> A[w] -> v = v;
			g -> A[w] -> next = NULL;
		} else {	
			struct list_node tmp;
			while(g -> A[w] -> next != NULL) {
				tmp = g -> A[w] -> next;
			}
			tmp = malloc(sizeof(list_node));
			tmp -> v = v;
			tmp -> next = NULL;
		}
		break;
	}
}

void graph_print(Graph g) {
	for(int i = 0; i < g -> n; i++) {
		struct list_node tmp = g -> A[i];
		printf("%d) ", i);
		printf("%d", g -> A[i] -> v);
		while(tmp -> next != NULL) {
			tmp = tmp -> next;
			printf(" -- %d", tmp -> v);
		}
		printf("\n");
	}
}
