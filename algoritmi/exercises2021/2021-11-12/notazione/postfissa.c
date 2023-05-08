#include <stdio.h>
#include <ctype.h>
#include <stdlib.h>
#include "../stack/stack.h"
#include "../item/item.h"

int main() {
    char token[10];
    Item op1, op2, ris;
    
    printf("Inserisci un numero o un operatore: ");
    while(scanf("%s", token) != EOF) {
        if(isdigit(token[0])) {
            push(atoitem(token));
        } else {
            op2 = pop();
            if(is_empty()) {
                if(token[0] == '+' || token[0] == '-') {
                    op1 = 0;
                } else {
                    op1 = 1;
                }
            } else {
                op1 = pop();
            }
            
            switch(token[0]) {
                case '+':
                    ris = op1 + op2;
                    break;
                case '-':
                    ris = op1 - op2;
                    break;
                case '*':
                    ris = op1 * op2;
                    break;
                case '/':
                    if(op2 == 0) {
                        printf("non e' possibile dividere per zero.\n");
                        exit(EXIT_FAILURE);
                    }
                    ris = op1 / op2;
                    break;
            }
            push(ris);
        }
    printf("Inserisci un numero o un operatore: ");
    }
    printf("\nRisultato = ");
    print_item(pop());
    printf("\n");
}	
