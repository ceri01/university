#include <stdio.h>
#include <string.h>
void farf(char *str, int len);

int main(int argc, char **argv) {
    int i = 1;
    while(i < argc) {
        farf(argv[i], strlen(argv[i]));
        i++; 
    }
}

void farf(char *str, int len) {
    char *p = str;

    while(p < str + len) {
        if (*p == 'a' || *p == 'e' || *p == 'i' || *p == 'o' || *p == 'u') {
            printf("%cf%c", *p, *p); 
        } else {
            printf("%c", *p);
        }
        p++;
    }
    printf("\n");
}
