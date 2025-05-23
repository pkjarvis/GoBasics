package main

import (
	"fmt"
	"time"
	"sync"
)


func say(s string){
	for i:=0;i<5;i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int){
	sum:=0 
	for _,v:=range s {
		sum+=v
	}
	c<-sum // send sum to c
}

func fibonacci(n int,c  chan int){
	x,y:=0,1
	for i:=0;i<n;i++{
		c<-x 
		x,y=y,x+y
	}
	close(c)
}

func fibonacci1(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

type SafeCounter struct{
	mu sync.Mutex
	v map[string]int
}

func (c*SafeCounter) Inc(key string){
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}
func (c*SafeCounter) Value(key string) int{
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}



func main(){
	go say("world")
	say("hello")

	s:=[]int{7,2,8,-9,4,0}

	c:=make(chan int)
	go sum(s[:len(s)/2],c)
	go sum(s[len(s)/2:],c)

	x,y:=<-c,<-c   // recieve from c

	fmt.Println(x,y,x+y)

	ch:=make(chan int,2)
	ch<-1
	ch<-2
	fmt.Println(<-ch)
	fmt.Println(<-ch)


	d:=make(chan int,10)

	go fibonacci(cap(d),d)
	
	for i:=range d{
		fmt.Println(i)
	}

	c1:=make(chan int)
	quit:=make(chan int)
	go func(){
		for i:=0;i<10;i++{
			fmt.Println(<-c1)
		}
		quit<-0
	}()

	fibonacci1(c1,quit)

	start:=time.Now()
	tick:=time.Tick(100*time.Millisecond)
	fmt.Println(start)
	fmt.Println(tick)
	boom:=time.After(500*time.Millisecond)
	elapsed:=func() time.Duration{
		return time.Since(start).Round(time.Millisecond)
	}
	for {
		select {
		case <-tick:
			fmt.Printf("[%6s] tick.\n",elapsed())
		case <-boom:
			fmt.Printf("[%6s] Boom!\n",elapsed())
			return 
		
		default:
			fmt.Printf("[%6s] .\n",elapsed())
			time.Sleep(50*time.Millisecond)

	    }
    }

	sc:=SafeCounter{v:make(map[string]int)}
	for i:=0;i<100;i++{
		go sc.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(sc.Value("somekey"))






	








}