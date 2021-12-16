#ifndef L_QUEUE_H
#define L_QUEUE_H
#include "b_list.h"
#include <stdbool.h>
#include <stdio.h>

typedef struct queue *Queue;

// Crea una coda
Queue create_queue(Queue q);

// Distrugge una coda
void destroy_queue(Queue *q);

// Accoda un elemento
void enqueue(Queue q, int i);

// Rimuove un elemento dalla coda
int dequeue(Queue q);

// Ritorna il primo valore della coda
int front_value(Queue q);

// Ritorna true se la coda e' vuota, false altrimenti
bool is_empty(Queue q);

// Stampa la coda
void print_queue(Queue q);

// Ritorna il numero degli elementi in codai (-1 se la coda non esiste)
int elements_in_queue(Queue q);

#endif

