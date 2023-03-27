#include <stdio.h>
#define N 3

typedef char Stringa[N];

typedef struct {
    Stringa content;
    int size;
} Paletto;

void hanoi(int n, Paletto from, Paletto tmp, Paletto to);

int main() {
    Paletto from;
    Paletto tmp;
    Paletto to;
    int n;
    char let = 'A';
    printf("Inserisci altezza pila: ");
    scanf(" %d", &n);
    printf("\n");

    from.size = n - 1; 
    tmp.size = 0;
    to.size = 0; 

    for(int i = 0; i < n; i++) {
        from.content[i] = let;
        tmp.content[i] = '\0';
        to.content[i] = '\0';
        let++;
    }

    //printf("%s %d", from.content, from.size);
    hanoi(n, from, tmp, to);
}

void hanoi(int n, Paletto from, Paletto tmp, Paletto to) {
    if (n == 1) {
        to.content[to.size] = from.content[from.size];
        from.content[from.size] = '\0';
        to.size++;
        from.size--;
        printf("from:%s\nto:%s\ntmp:%s\n", from.content, tmp.content, to.content);
        printf("\n");
        return;
    }
    hanoi(n-1, from, to, tmp); 
    to.content[to.size] = from.content[from.size];
    to.size++;
    from.size--;
    printf("from:%s\nto:%s\ntmp:%s\n", from.content, tmp.content, to.content);
    printf("\n");
 
    hanoi(n-1, tmp, from, to);
}
