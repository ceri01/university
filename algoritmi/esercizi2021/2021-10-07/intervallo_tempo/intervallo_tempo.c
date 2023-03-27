#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

void split(char *str, int *arr, size_t length) {
    char tmp_str[3];
    
    for (int i = 0, arr_i = 0; i < length; i++) {
        if (str[i] == ':' || str[i] == '\0') { 
            for (int k = (i-2); k < i; k++) {
                strncat(tmp_str, &str[k], 1);
            }
            arr[arr_i] = atoi(tmp_str);
            arr_i++;
            tmp_str[0] = '\0';
        }
    }
}

void time_distance(int *ris, int *arr1, int *arr2, size_t length) {
    for (int i = 0; i < length; i++) {
        ris[i] = abs(arr1[i] - arr2[i]);
    }
}

int main() {
    char orario1[9];
    char orario2[9];
    int arr1[3];
    int arr2[3];
    int ris[3];

    printf("Inserisci primo orario (formato hh:mm:ss): ");
    scanf(" %s", orario1);
    printf("Inserisci secondo orario (formato hh:mm:ss): ");
    scanf(" %s", orario2);

    split(orario1, arr1, 9);
    split(orario2, arr2, 9);
    time_distance(ris, arr1, arr2, 3);

    printf("La distanza tra i due orari è di %d ore, %d minuti e %d secondi.\n", ris[0], ris[1], ris[2]);
}

