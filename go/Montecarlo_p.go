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
			go func(){
				ch <- 1
				fmt.Println(len(ch))
				loop_end := lim / 100
//				if len(ch) != thread_count - 1{
//					loop_end := lim/(thread_count - 1)
//				} else {
//					loop_end := lim % (thread_count - 1)
//				}
				fmt.Println(len(ch), loop_end)
				for j := 0; j < loop_end; j++{
					x := random.Float64()
					y := random.Float64()
					if x*x + y*y <= 1{
						mutex.Lock()
						count += 1
						mutex.Unlock()
					}
				}
			}()

/*		} else {
			fmt.Println(lim % thread_count)
			for j := 0; j < lim % thread_count; j++{
				x := random.Float64()
				y := random.Float64()
				if x*x + y*y <= 1{
					mutex.Lock()
					count += 1
					mutex.Unlock()
				}
			}
		}
*/
	}
	for i := 0; i < thread_count; i++{
		a := 0
		a = <- ch
		a++
	}
/*	limf := 2147483647.0
	pi := count / limf
	fmt.Println(pi)
*/
}
