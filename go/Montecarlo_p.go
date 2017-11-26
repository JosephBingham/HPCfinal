package main


import ("fmt"
	"math/rand"
	"time"
	"sync/atomic"
	)

func throwDart(num int, count *int, mutex *sync.Mutex){
	for i := 0; i < num; i++ {
		x := random.Float64()
		y := random.Float64()
		if x*x + y*y <= 1 {
			mutex.Lock()
			*count++
			mutex.Unlock()
		}
	}
}


func main(){
	var mutex = &sync.Mutex{}
	thread_count := 4
	count := 0
	lim := 2147483647
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	for i := 0; i < thread_count - 1; i++ {
		go throwDart((lim/(thread_count - 1)), &count, &mutex)
	}
	go throwDart(lim % thread_count, &count);
	countf := 1.0 * count
	limf := 2147483647.0
	pi := 4.0 * countf / limf
	fmt.Println(pi)
}
