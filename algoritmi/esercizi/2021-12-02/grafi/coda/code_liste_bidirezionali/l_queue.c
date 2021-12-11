#include <stdbool.h>
#include <stdlib.h>
#include <stdio.h>
#include "l_queue.h"

struct queue {
    Node head;
    Node tail;
    bool empty;
    int elements;    
};

typedef struct queue *Queue;

Queue create_queue(Queue q) {
    if(q == NULL) {
        q = malloc(sizeof(struct queue));
        q -> head = NULL;
        q -> tail = NULL;
        q -> empty = true;
        q -> elements = 0;   
    }
    return q;
}

void destroy_queue(Queue *q) {
    Queue tmp = *q;
    if(tmp -> head == tmp -> tail) {
        free(tmp -> head);
    } else if(tmp -> head != NULL) {
        while(tmp -> tail -> prev != NULL) {
            tmp -> tail = tmp -> tail -> prev;
            free(tmp -> tail -> next);
        }
        free(tmp -> tail);       
    }
    free(tmp);
    *q = NULL;
}

void enqueue(Queue q, int i) {
    if(q != NULL) {
        Node new_node = malloc(sizeof(struct node));
        new_node -> info = i;
        if(q -> head == NULL) {
            new_node -> next = NULL;
            new_node -> prev = NULL;
            q -> head = new_node;
            q -> tail = new_node;
        } else {
            q -> tail -> next = new_node;
            new_node -> next = NULL;
            new_node -> prev = q -> tail;
            q -> tail = new_node;
        }
        q -> empty = false;
        (q -> elements)++;
    }
}

int dequeue(Queue q) {
    if(q -> head == NULL) {
        printf("La coda e' vuota, impossibile torgliere elenenti, valore ritornato = NULL\n");
    } else {
        int tmp = q -> head -> info;
        if(q -> head == q -> tail) {
            free(q -> head);
            q -> head = NULL;
            q -> tail = NULL;
            q -> empty = true;
            q -> elements--;
            return tmp;
        } else {
            Node new_head = q -> head -> next;
            free(q -> head);
            q -> head = new_head;
            q -> head -> prev = NULL;
            q -> elements--;
            return tmp;
        }
    }
    return NULL;
}

int front_value(Queue q) {
    if(q != NULL) {
        return q -> head -> info;
    } else {
        printf("Errore, la coda e' vuota, il valore tornato e' NULL\n");
    }
    return NULL; 
}

bool is_empty(Queue q) {
    if(q != NULL) {  
        return q -> empty;
    }
    return false;
}

int elements_in_queue(Queue q) {
    if(q != NULL) {  
        return q -> elements;
    }
    return -1;
}


void print_queue(Queue q) {
    if(q != NULL) {
        if(q -> head != NULL) {
            Node tmp = q -> head;
            while(tmp != NULL) {
                print_item(tmp -> info);
                tmp = tmp -> next;
            }
            printf("\n");
            free(tmp);       
        }
    }
}

