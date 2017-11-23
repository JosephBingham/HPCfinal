import java.util.*;

public class Montecarlo_s{


    public static void main(String[] args){
	Random rand = new Random();
	double pih = 0;
	int[] vals = new int[2];
	int lim = Integer.MAX_VALUE;
	double x;
	double y;
	for(int i = 0 ; i < lim; i++){
		x = rand.nextDouble();
		y = rand.nextDouble();
		if(x*x + y*y < 1){
			vals[0]++;
	    	}
	    	else{
			vals[1]++;
	    	}
	}
	pih = 4.0 * (vals[0] + 0.0)/(vals[1] + vals[0]);
	System.out.println(pih);

    }

}
