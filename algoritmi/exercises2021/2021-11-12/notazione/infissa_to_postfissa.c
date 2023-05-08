#include <stdlib.h>
#include <stdio.h>
#include <ctype.h>
#include <string.h>
#include "../item/item.h"
#include "../stack/stack.h"

int main() {
    char ch;
    while((ch = getchar()) != '\n') {
        char *str = malloc(2 * sizeof(char));
        str[1] = '\0';
        if(isdigit(ch)) {
            *str = ch;
            print_item(atoitem(str));
            printf(" ");
        } else {
            switch(ch) {
                case '/':
                case '*':
                case '-':
                case '+':
                    *str = ch;
                    push(atoitem(str));
                    break;
                case ')':
                    str = pop();
                    print_item(atoitem(str));
                    printf(" ");
                    free(str);
                    break;
            }
        }
    } 
}






