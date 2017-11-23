package main

import "fmt"
import "math/rand"
import "time"

func main() {
	val := 0.0
	lim := 100000000
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	for i := 0; i < lim; i++ {
		x := random.Float64()
		y := random.Float64()
		if x*x + y*y <= 1 {
			val++
		}
		i++
	}
	limf := 100000000.0
	fmt.Println((val*1.0)/limf)
	end_val := (4.0 * (val*1.0)/limf)
	fmt.Println(end_val)
}
