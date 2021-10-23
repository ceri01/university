#include <stdio.h>
#include <math.h>
#include <ctype.h>
#define L 100

void print_array(int a[L], int len); 

void insertion_sort(int a[L], int len, int val); 

int main() {
    int arr[L], act_len = 0, n = -1;
    
    while(act_len < L) { 
        scanf("%d", &n);
        if (n == 0) {
            break;
        }
        insertion_sort(arr, act_len + 1, n);
        print_array(arr, act_len + 1);
        act_len++;
    }
    printf("Array finale:\n");
    print_array(arr, act_len);
}

void print_array(int a[L], int len) {
    for(int i = 0; i < len; i++) {
        printf("%d ", a[i]);
    }
    printf("\n");
}

void insertion_sort(int a[L], int len, int val) {
    int x, find;
    
    a[len-1] = val;
    for(int index = 1; index < len; index++) {
        x = a[index];
        find = index-1;
        while(find >= 0 && a[find] > x) {
            a[find+1] = a[find];
            find = find - 1;
        }
        a[find+1] = x;
    }
}
