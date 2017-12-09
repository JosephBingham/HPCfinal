package main

import "fmt"
import "math/rand"
import "time"
import "sync/atomic"

func main() {
	thread_count := 4
	var in int64 = 0
	var total int64 = 0
	lim := 2147483647
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	ch := make(chan int)
	for i := 0; i < thread_count; i++{
		go func(){
			ch <- 1
			for i := 0; i < lim/thread_count; i++ {
				x := random.Float64()
				y := random.Float64()
				if x*x + y*y <= 1 {
					atomic.AddInt64(&in, 1)
				}
				atomic.AddInt64(&total, 1)
			}
		}()
	}
	for i := 0; i < thread_count; i++{
		<-ch
	}
	countf := float64(in)
	totalf := float64(total)
	pi := (4.0 * countf/totalf)
	fmt.Println(pi)
}
