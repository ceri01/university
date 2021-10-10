#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

// funzione che permette di acquisire in input una stringa allocando la memoria

char * get_word(int size, int * pt_size) {
    char *word = NULL, ch;
    int tmp_size = 0;

    word = malloc(size);
    
    while((ch = getchar()) != EOF) {
        // printf("tmp_size %d\n", tmp_size); 
        if (ch == '\n') {
            ch = '\0';
            break;
        }
        if (size <= tmp_size) {
            size += 3;
            word = realloc(word, size);
        }
        word[tmp_size++] = ch;
        
    }
    *pt_size = size;
    return word;
}

// Funzione che prende la stringa con l'orario in formato am/pm, prende il valore di ora e minuti (li converte in int) e li mette in un array

void split_am_pm(char *str, int *arr, size_t length) {
    char tmp_str[3] = "";
    
    for (int i = 0, arr_i = 0; i < length; i++) {
        if (str[i] == ':' || str[i] == ' ') { 
            for (int k = (i-2); k < i; k++) {
                strncat(tmp_str, &str[k], 1);
            }
            arr[arr_i] = atoi(tmp_str);
            arr_i++;
            tmp_str[0] = '\0';
        }
        if (str[i] == ' ') {
            break;
        }
    }
}

// Funzione che prende la stringa con l'orario normale, prende il valore di ora e minuti (li converte in int) e li mette in un array

void split_normal(char *str, int *arr, size_t length) {
    char tmp_str[3] = "";
    
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


// Funzione che coverte un orario in un orario in formato am/pm

void n_to_am_pm(int *arr) {
    if (arr[0] > 12) {
        arr[0] = arr[0] - 12;
    } else if (arr[0] == 0) {
        arr[0] = arr[0] + 12;
    }
}

// Funzione che coverte un orario nel formato am/pm in un orario normale

void am_pm_to_n(int *arr, char *am_pm) {
    if ((strcmp(am_pm, "pm") == 0) && (arr[0] != 12)) {
        arr[0] = arr[0] + 12;
    } else if ((strcmp(am_pm, "am") == 0) && (arr[0] == 12)) {
        arr[0] = arr[0] - 12;
    } 
}

// Funzione che setta la stringa am ad "am" oppure "pm" in base all'orario normale inserito

void set_am_pm(char *str, char *am_pm) {
   if (str[6] == 'a') {
       am_pm[0] = 'a';
   } else {
       am_pm[0] = 'p';   
   }
   am_pm[1] = 'm';
}

// Funzione che setta la stringa am ad "am" oppure "pm" in base all'orario nel formato am/pm inserito

void set_am_pm_res(int *arr, char *am_pm) {
   if ((arr[0] >= 0 && arr[0] < 12)) {
       am_pm[0] = 'a';
   } else {
       am_pm[0] = 'p';   
   }
   am_pm[1] = 'm';
}


int main() {
    char *input_hours;
    int size = 6;
    int *pt_size = &size;
    char am[2];
    int hrs[2];

    printf("Inserisci orario (formato hh:mm oppure hh:mm am/pm): ");
    input_hours = get_word(size, pt_size);

    if (size > 6) {
        split_am_pm(input_hours, hrs, size);
        // printf("%d", hrs[0]);
        if (hrs[0] > 12 || hrs[0] < 1) {
            printf("Errore inserimento orario.\n");
            return 0;
        }
        set_am_pm(input_hours, am);
        am_pm_to_n(hrs, am);
        printf("Orario inserito = %s\nOrario convertito = %d:%d\n", input_hours, hrs[0], hrs[1]);
    } else {
        split_normal(input_hours, hrs, size);
        // printf("%d", hrs[0]);
        if (hrs[0] > 23 || hrs[0] < 0) {
            printf("Errore inserimento orario.\n");
            return 0;           
        }
        set_am_pm_res(hrs, am);
        n_to_am_pm(hrs);
        printf("Orario inserito = %s\nOrario convertito = %d:%d %s\n", input_hours, hrs[0], hrs[1], am);
    }
}

