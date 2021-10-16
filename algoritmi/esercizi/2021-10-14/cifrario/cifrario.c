#include <stdio.h>
#include <string.h>

int main() {
    char c, str[100];
    int key, index = 0;

    while ((c = getchar()) != '.') {
        str[index] = c;
        index++;    
    }
    str[index] = c;
    index = 0;
    
    printf("Inserire la chiave di cifratura (numero intero): \n");
    scanf(" %d", &key);

    for (; index < strlen(str); index++) {
        str[index] = str[index] + key;
    }
    printf("%s\n", str);
}
