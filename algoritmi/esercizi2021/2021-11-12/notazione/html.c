#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdbool.h>
#include "../stack/stack.h"
#include "../item/item.h"
#include "../list/list.h"

bool compare(char *str);

int main(void) {
    char *tag;
    char ch;
    int dim = 3;
    int i = 0;

    while(ch != '\n') {
        tag = malloc(dim * sizeof(char));
        while((ch = getchar()) != ' ' && ch != '\n') {    
            if(tag == NULL) {
                printf("Error!\n");
                exit(EXIT_FAILURE);
            }
            if(i >= dim) {
                dim += 1;
                tag = realloc(tag, dim);
            }
            tag[i] = ch;
            i++;
        }
    
        if(i >= dim) {
            dim += 1;
            tag = realloc(tag, dim);
        } 
        tag[i++] = '\0';
        switch(*(tag + 1)) {
            case '/':
                if(compare(tag)) {
                    pop();
                    free(tag);
                }
                break;
            default:
                push(tag);
                break;
        }
        dim = 3;
        i = 0;
    }
    if(is_empty()) {
        print_stack();
        printf("Fatto\n");
    } else {
        print_stack();
        printf("documento non ben formato\n");
    }
}

bool compare(char *str) {
    int dim1 = strlen(top()); // apertura
    int dim2 = strlen(str); // chiusura
    if((dim2 - dim1) != 1) {
        printf("Documento non ben formato!\n");
        exit(EXIT_FAILURE);
    }
    char *sub1 = malloc(sizeof(char) * (dim1 - 1));
    char *sub2 = malloc(sizeof(char) * (dim2 - 2));
    sub1[dim1-2] = '\0';
    sub2[dim2-3] = '\0';
    char *p = top();
    memcpy(sub1, &p[1],(dim1 - 2));
    memcpy(sub2, &str[2], (dim2 - 3)); 
    if(strcmp(sub1, sub2) != 0) {
        printf("Documento non ben formato!\n");
        exit(EXIT_FAILURE);
    }
    free(sub1);
    free(sub2);
    return true;
}

