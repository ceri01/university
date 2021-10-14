#include <stdio.h>
#include <ctype.h>
int main() {
    int c, f[10] = {0};
    c = getchar();
    while(c != '.') {
        if (isdigit(c)) {
           f[c - '0']++;
        }
        c = getchar();
    }
    printf("Valori inseriti:\n");
   // valori inseriti
   //
   // for(int i = 0; i < 10; i++) {
   //   if (f[i] == 0) {
   //       continue;
   //   }
   //   printf("%d ", i);
   // }
   

   // valori ripetuti
   for (int i = 0; i < 10; i++) {
       if (f[i] > 1) {
            printf("%d ", i);
       }
   }
   printf("\n");
}
