#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>
#include "queue.h"
#include "../../2021-11-12/item/item.h"
#define N 10

struct queue {
    int head;
    int rear;
    int elements;
    bool is_full;
	bool is_empty;
    int *qu;
};

Queue *create_queue(Queue *q) {
	q = malloc(sizeof(Queue));
	if(q == NULL) {
		exit(EXIT_FAILURE);
	}
	q -> elements = 0;
	q -> head = 0;
	q -> rear = N - 1;
	q -> is_full = false;
	q -> is_empty = true;
	q -> qu = malloc(N * sizeof(Item));
	return q;
}

void destroy_queue(Queue *q) {
	free(q -> qu);
	free(q);
}

void enqueue(Queue *q, Item i) {	
	if(q -> is_full) {
		printf("Queue is full\n");
		return;
	}
	
	q -> rear = (((q -> rear) + 1) % N);
	q -> is_empty = false;
	(q -> elements)++;
	q -> qu[q -> rear] = i;
	if(q -> is_empty == false && ((((q -> rear) + 1) % N) == q -> head)) {
		q -> is_full = true;
	}
}

Item dequeue(Queue *q) {
	if(q -> is_empty) {
		printf("Queue is empty\n");
		return;
	}

	q -> is_full = false;
	Item val = front_value(q);
	(q -> elements)--;
	q -> head = (((q -> head) + 1) % N);
	if(q -> is_full == false && ((((q -> rear) + 1) % N) == q -> head)) {
		q -> is_empty = true;
	}
	return val;
}

Item front_value(Queue *q) {
	return q -> qu[q -> head];
}

void print_queue(Queue *q) {
	if(q -> is_empty) {
		printf("\n");
		return;
	}
	int i;
	for(i = q -> head; i != q -> rear;) {
		print_item(q -> qu[i]);
		printf(" ");
		i = ((i + 1) % N);
	}
	print_item(q -> qu[i]);
	printf("\n");
}

bool is_full_queue(Queue *q) {
	return q -> is_full;
}

bool is_empty_queue(Queue *q) {
	return q -> is_empty;
}
