package main

import "fmt"

func f( x float64) float64{
	sum := 0.0
	for i := 0; i < int(x)/10; i++{
		sum += float64(i)
	}
	return sum
}

func trap( a float64, b float64, h float64, sum *float64) {
	*sum += (f(a)+f(b))*h
}

func main(){
	var sum float64 = 0
	n := 200.0
	h := 1000000000.0/n
	var mesh [201]float64
	for i := 0; i <= int(n); i++{
		mesh[i] = h*float64(i)
	}
	for i := 0; i < int(n); i++{
		trap(mesh[i],mesh[i+1],h, &sum)
	}
	fmt.Println(sum)
}
