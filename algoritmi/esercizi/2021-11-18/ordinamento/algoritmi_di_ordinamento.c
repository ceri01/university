#include <stdio.h>
#include <stdlib.h>

void print(int a[], int length);
void swap(int a[], int i, int j);
void selection_sort(int a[], int n);
void mergesort(int a[], int sx, int dx, int tmp[]);
void merge_sort(int a[], int dim);
void merge(int a[], int sx, int m, int dx, int tmp[]);
    
int main() {
	int n, d1 = 0, d2 = 0;

	printf("Dimensione primo array: ");
	scanf("%d", &d1);
	int a[d1];

	printf("Inserisci i valori per il primo array:\n");
	for(int i = 0; i < d1; i++) {
		scanf("%d", &n);
		a[i] = n;
	}

	printf("Dimensione secondo array: ");
	scanf("%d", &d2);	
	int b[d2];

	printf("Inserisci i valori per il secondo array:\n");
	for(int i = 0; i < d2; i++) {
		scanf("%d", &n);
		b[i] = n;
	}
	printf("\n");
	print(a, d1);
	selection_sort(a, d1);
	printf("Selection Sort = ");
	print(a, d1);
	printf("\n");
	print(b, d2);
	merge_sort(b, d2);
	printf("Merge Sort = ");
	print(b, d2);
}

void print(int a[], int length) {
	for(int i = 0; i < length; i++) {
		printf("%d ", a[i]);
	}
	printf("\n");
}

void swap(int a[], int i, int j) {
	int tmp = a[i];
	a[i] = a[j];
	a[j] = tmp;
}

void selection_sort(int a[], int n) {
	if(n <= 1) {
		return;
	}
	int max_index = 0;
	for(int i = 1; i < n; i++) {
		if(a[max_index] < a[i]) {
			max_index = i;
		}
	}
	swap(a, max_index, n-1);
	selection_sort(a, n-1);
}

void merge_sort(int a[], int dim) {
    int tmp[dim];
    mergesort(a, 0, dim, tmp);
}

void mergesort(int a[], int sx, int dx, int tmp[]) {
	int m = (dx + sx) / 2;

    if(dx-sx <= 1) {
        return;
	}

	mergesort(a, sx, m, tmp);	
	mergesort(a, m, dx, tmp);
    	merge(a, sx, m, dx, tmp);   
}

void merge(int a[], int sx, int m, int dx, int tmp[]) {
	int index = 0;
    int i = sx, k = m;

    while(i < m && k < dx) {
		if(a[i] <= a[k]) {
			tmp[index] = a[i];
			i++;
		} else {
			tmp[index] = a[k];
			k++;
		}
		index++;
	}

	while(i < m) {
		tmp[index] = a[i];
		i++;
		index++;
	}

	while(k < dx) {
		tmp[index] = a[k];
		k++;
		index++;
	}

    i = 0;
    while(i < dx-1-sx) {
        a[i+sx] = tmp[i];
        i++;
    }
}

