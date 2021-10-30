#include <stdio.h>
#include <ctype.h>
#include <string.h>

char *upper(char *str);

int main() {
    char str[100], *str_ptr = str;
    char ch;

    while((ch = getchar()) != '\n') {
        *str_ptr = ch;
        str_ptr++;
    }
    str_ptr = '\0';
    printf("%s\n", upper(str));
}

char *upper(char *str) {
    char *ptr_first = str;
    while(*str != '\0') {
        *str = toupper(*str);
        str++;
    }
    return ptr_first;
}
