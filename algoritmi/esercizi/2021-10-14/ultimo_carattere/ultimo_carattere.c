#include <stdio.h>
#include <string.h>

int remove_char(char *s, int i, int len) {
    int init_ind = i;
    if (s[i-2] == ' ') {
        s[i-1] = '\\';
        return i+1;
    }

    while (i <= len || s[i] != '\0') {
        s[i-1] = s[i];
        i++;
    }
    return init_ind;
}



int main() {
    char c, str[100];
    int index = 0, len;
    
    while ((c = getchar()) != '\n') {
        if (index >= 98) {
            index++;
            break;
        }
        str[index] = c;
        index++;
    }
    str[index] = '\0';
    len = strlen(str);

    for (index = 0; index < len; index++) {
        if (str[index] == ' ') {
            index = remove_char(str, index, len);
        }
    }
    str[strlen(str)-1] = '\0';
    len = strlen(str);

    printf("%s\n", str);
}
