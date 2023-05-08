#ifndef CIRC_QUEUE_H
#define CIRC_QUEUE_H
#include "../../2021-11-12/item/item.h"
#include <stdbool.h>

typedef struct queue Queue;

// Create queue
Queue *create_queue(Queue *q);

// Destroy queue
void destroy_queue(Queue *q);

// Put element on rear
void enqueue(Queue *q, Item i);

// Remove and return element from front
Item dequeue(Queue *q);

// Return front element
Item front_value(Queue *q);

// Check if queue is empty
bool is_empty_queue(Queue *q);

// Print queue content
void print_queue(Queue *q);

bool is_full_queue(Queue *q);

#endif
