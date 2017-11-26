package main

import "fmt"
import "math/rand"
import "time"

func main() {
	count := 0.0
	lim := 2147483647
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	for i := 0; i < lim; i++ {
		x := random.Float64()
		y := random.Float64()
		if x*x + y*y <= 1 {
			count++
		}
	}
	limf := 2147483647.0
	pi := (4.0 * (count*1.0)/limf)
	fmt.Println(pi)
}
