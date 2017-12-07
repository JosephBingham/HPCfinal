import java.io.*;
import java.util.*;
public class PHS_s{

	public static int[] PHS_sort(int[] arr){
		if(arr.length < 2){
			return arr;
		}
		int max = arr[0];
		int min = arr[0];
		for( int i : arr){
			if( i < min ){
				min = i;
			}
			if( i > max ){
				max = i;
			}
		}
		int[] temp = new int[max-min+1];
		for( int i : arr){
			temp[i-min] = i;
		}
		int j = 0;
		for( int i : temp){
			if( i != 0 ){
				arr[j] = i;
				j++;
			}
		}
		return arr;
	}
	private static void shuffleArray(int[] array){
		int index, temp;
		Random random = new Random();
		for (int i = array.length - 1; i > 0; i--){
			index = random.nextInt(i + 1);
			temp = array[index];
			array[index] = array[i];
			array[i] = temp;
		}
	}
	public static void main(String[] args){
		int[] arr = new int[214748364];
		for(int i = 0; i < 214748364; i++){
			arr[i] = i + 1;
		}
		shuffleArray(arr);
		int[] result = PHS_sort(arr);
		for( int i : result ){
			System.out.println(i);
		}
	}
}
