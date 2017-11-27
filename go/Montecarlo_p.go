package main


import ("fmt"
	"math/rand"
	"time"
	"sync"
	)


func main(){
	var mutex = &sync.Mutex{}
	ch := make(chan int)
	thread_count := 4
	count := 0.0
	lim := 2147483647
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	for i := 0; i < thread_count; i++ {
		if i != thread_count -1 {
			go func(){
				ch <- 1
				for i := 0; i < lim/(thread_count - 1); i++{
					x := random.Float64()
					y := random.Float64()
					if x*x + y*y <= 1{
						mutex.Lock()
						count += 1
						mutex.Unlock()
					}
				}
			}()

		} else {
			for i := 0; i < lim % thread_count; i++{
				x := random.Float64()
				y := random.Float64()
				if x*x + y*y <= 1{
					mutex.Lock()
					count += 1
					mutex.Unlock()
				}
			}
		}

	}
	for i := 0; i < thread_count - 1; i++{
		a := 0
		a <- ch
		a++
	}
	limf := 2147483647.0
	pi := count / limf
	fmt.Println(pi)
}
