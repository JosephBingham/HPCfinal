#include<stdio.h>
#include<stdlib.h>
#include<omp.h>

int *PHS(int *arr){
	int max = arr[0];
	int min = arr[0];
	#pragma omp parallel for num_threads(16)
	for(int i = 0; i < 214748364; i++){
		if( arr[i] < min ){
			#pragma omp critical
			min = arr[i];
		}
		if( arr[i] > max ){
			#pragma omp critical
			max = arr[i];
		}
	}
	int *temp = (int *)calloc((max - min + 1), sizeof(int));
	#pragma omp parallel for num_threads(16)
	for( int i = 0; i < (max - min + 1); i++){
		temp[i-min] = i;
	}
	int j = 0;
	#pragma omp parallel for num_threads(4)
	for( int i = 0; i < (max - min + 1); i++){
		if( i != 0 ){
			#pragma omp critical
			{
			arr[j] = i;
			j++;
			}
		}
	}
	free(temp);
	return arr;
}
void shuffleArray(int *arr){
	int index;
	int temp;
	for(int i = 214748364; i > 0; i--){
		index = (rand() % i) + 1;
		temp = arr[index];
		arr[index] = arr[i];
		arr[i] = temp;
	}
}

int main(void){
	int *arr = (int *)malloc(214748364*sizeof(int));
	#pragma omp parallel for num_threads(4)
	for(int i = 1; i <= 214748364; i++){
		arr[i] = i;
	}
	shuffleArray(arr);
	arr = PHS(arr);
	for( int i = 0; i < 214748364; i++){
		printf("%d\n", arr[i]);
	}
	return 0;
}
