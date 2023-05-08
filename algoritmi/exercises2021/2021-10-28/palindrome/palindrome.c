#include <stdio.h>
#include <string.h>
#include <stdbool.h>

bool palin(char *s, int len);

int main(int argc, char **argv) {
    int len = 0;
    int index = 1;
    char *str;
    
    while (index < argc) {
        str = argv[index];

        for(char *p = argv[index]; *p != '\0'; p++){ 
            len++;
        }
        
        if(palin(str, len-1)) {
            printf("%s e'palindroma\n", str);
        } else {
            printf("%s non e' palindroma\n", str);
        }
        len = 0;
        index++;
    }
}

bool palin(char *s, int len) {
    char *i = s;
    char *f = s + len;
    while(i <= f ) {
        if(*i != *f) {
            return false;
        }
        i++;
        f--;
    }
    return true;
}
