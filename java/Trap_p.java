public class Trap_s{

	private static double f(double x){
		int sum = 0;
		int n = (int) x;
		for(int i = 0; i < n/10; i++){
			sum += i;
		}
		return 1.0*sum;
	}

	private static double trap(double a, double b, double h){
		return (f(a)+f(b))*h;
	}

	public static void main(String[] args){
		double sum = 0.0;
		int n = 200;
		double h = 1000000000/n;
		double[] mesh = new double[n+1];
		// omp parallel for numThreads(16)
		for(int i = 0; i <= n; i++){
			mesh[i] = h*i;
		}
		// omp parallel for numThreads(16)
		for(int i = 0; i < n; i++){
			// omp critical
			sum += trap(mesh[i], mesh[i+1], h);
		}
		System.out.println(sum);
	}

}
