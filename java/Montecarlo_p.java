import java.util.*;

public class Montecarlo_p{


    public static void main(String[] args){
	Random rand = new Random();
	double pih = 0;
	int count = 0;
	int lim = Integer.MAX_VALUE;
	double x;
	double y;
	// omp parallel for threadNum(16)
	for(int i = 0 ; i < lim; i++){
		x = rand.nextDouble();
		y = rand.nextDouble();
		if(x*x + y*y < 1){
			// omp critical
			count++;
	    	}
	}
	pih = 4.0 * (count + 0.0)/lim;
	System.out.println(pih);

    }

}
