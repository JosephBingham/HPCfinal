
#include<stdio.h>
#include<stdlib.h>
#include<omp.h>
double random_f(void){
      double r = (double)rand()/(double)RAND_MAX;
      return r;
}

int main(void){
	double count = 0.0;
	double pih = 0;
	int lim = 2147483647;
	double dlim = 2147483647.0;
	double x;
	double y;
	int i;
#pragma omp parallel for num_threads(16)
	for(i = 0; i < lim; i++){
		x = random_f();
		y = random_f();
		if(x*x + y*y < 1){
#pragma omp critical
			count++;
		}
	}
	pih = 4.0 * count/dlim;
	printf("%lf\n", pih);
}


