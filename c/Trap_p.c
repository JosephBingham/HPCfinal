#include<stdio.h>
#include<stdlib.h>
#include<unistd.h>
#include<omp.h>

double f( double x){
	int sum = 0;
	for(int j = 0; j < x/10; j++){
		sum += j;
	}
	return 1.0 * sum;
}


double trap(double a, double b, double h){
	return (f(a)+f(b))*h;
}

int main(void){
	long double sum = 0;
	int n = 200;
	double *mesh = (double *)malloc((n+1)*sizeof(double));
	double h = 1000000000/n;
	#pragma omp parallel for num_threads(16)
	for(int i = 0; i <= n; i++){
		mesh[i] = h*i;
	}
	#pragma omp parallel for num_threads(16)
	for(int i = 0; i < n; i++){
		sum += trap(mesh[i],mesh[i+1], h);
	}
	printf("%Lf\n",sum);
}
