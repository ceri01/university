#ifndef STACK_M
#define STACK_M

/* Svuota la pila */
void make_empty(void);

/* Restituisce 1 se la pila e' vuota 0 altrimenti */
int is_empty(void);

/* Restituisce 1 se la pila e' piena, 0 altrimenti*/
int is_full(void);

/* Se la pila non e' vuota restituisce il top della pila, altrimenti esce un messaggio di errore */
int top(void);

/* Se la pila non e' vuota estrae il top della pila, altrimenti esce un messaggio di errore */
int pop(void);

/* Se c' e' spazio, aggiunge n alla pila, altrimenti esce con messaggio di errore */
void push(int n);

/* Stampa il contenuto della pila, partendo da top */
void print_stack(void);

#endif
