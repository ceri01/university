#include <stdio.h>

int main() {
    int digit[10] = {0}, tot = 0, i;
    long n;

    printf("Inserisci numero: ");
    scanf(" %ld", &n);
    
    // k confronti dove k e' il numero di cifre del numero inserito (con un massimo di )
    for(; n >  0;) { 
        digit[n % 10]++;
        n /= 10;
    } 

    // 10 iterazioni -> al piu' 31 confronti -> O(1)
    for (i = 0; i < 10; i++) {
        tot = tot + digit[i]; 
        if (tot > (i+1)) {
            printf("Cifre ripetute: ");
            break;
        } else if (i == 9 && tot <= 10) {
            printf("Non ci sono cifre ripetute.\n");
            return 0;
        }
    }

    // 10 iterazioni quindi 20 confronti -> O(1)
    for (; i < 10; i++) {
        if (digit[i] > 1) {
            printf("%d ", i);
        }
    }
    printf("\n");
}
