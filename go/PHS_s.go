package main

import ("fmt"
	"math/rand"
	)

func PHS( arr []int) []int{
	if len(arr) < 2 {
		return arr
	}
	max := arr[0]
	min := arr[0]
	for i := 0; i < len(arr); i++{
		a := arr[i]
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
	}
	var temp = make([]int, max - min + 1)
	for _, i := range arr {
		temp[i-min] = i
	}
	j := 0
	for _, i := range temp {
		if i != 0 {
			arr[j] = i
			j++
		}
	}
	return arr
}

func main(){
	arr := rand.Perm(214748363)
	for i, _ := range arr {
		arr[i]++
	}
	result := PHS(arr)
	for i, _ := range result {
		fmt.Println(result[i])
	}
}
