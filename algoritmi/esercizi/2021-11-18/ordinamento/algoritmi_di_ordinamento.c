#include <stdio.h>
#include <stdlib.h>

void print(int a[], int length);
void swap(int a[], int i, int j);
void selection_sort(int a[], int n);
void merge_sort(int a[], int sx, int dx);

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
	merge_sort(b, 0, d2);
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

void merge_sort(int a[], int sx, int dx) {
	int *tmp1 = malloc(dx/2 * sizeof(int));
	int *tmp2 = malloc(dx/2 * sizeof(int));
	int dim_disp = 0;

	if(dx <= 1) {
		return;
	}

	for(int i = 0; i < dx/2; i++) {
		tmp1[i] = a[i];
	}

	if(dx % 2 == 1 && dx != 1) {
		dim_disp = (dx/2)+1;
		tmp2 = realloc(tmp2, dim_disp);
	} else {
		dim_disp = dx/2;				
	}

	for(int i = 0, k = dx/2; k < dx; i++, k++) {
		tmp2[i] = a[k];
	}

	merge_sort(tmp1, 0, dx/2);	
	merge_sort(tmp2, 0, dim_disp);

	int i = 0, k = 0, j = 0;
	while(j < dx/2 && k < dim_disp) {
		if(tmp1[j] <= tmp2[k]) {
			a[i] = tmp1[j];
			j++;
		} else if(tmp2[k] < tmp1[j]) {
			a[i] = tmp2[k];
			k++;
		}
		i++;
	}

	while(j < dx/2) {
		a[i] = tmp1[j];
		i++;
		j++;
	}
	while(k < dim_disp) {
		a[i] = tmp2[k];
		i++;
		k++;
	}
	free(tmp1);
	free(tmp2);
}


