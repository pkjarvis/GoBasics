package main

import (
	"fmt"
	// "unicode/utf8"
	"math"
	"errors"
	"time"
)

// Closures
func intSeq() func() int{
	i:=0
	return func() int{
		i++
		return i
	}
}

// Recursion
func fact(n int) int{
	if n==0{
		return 1
	}
	return n*fact(n-1)
}

// Pointers
func zeroval(ival int) {
	ival = 0
}
func zeroptr(iptr *int) {
	*iptr = 0
}

// Struct
type person struct{
	name string
	age int
}

func newPerson(name string)*person{
	p:=person{name:name}
	p.age=42
	return &p
}

// // Methods
// type rect struct{
// 	width,height int
// }

// func (r *rect) area() int{  // here in parameter we could pass both pointer as well as value
// 	return r.width*r.height
// }

// func (r rect) perim() int{
// 	return 2*(r.width)+2*(r.height)
// }

// Interface
type geometry interface{
	area() float64
	perim() float64
}

type rect struct{
	width,height float64
}

type circle struct{
	radius float64
}

func (r rect) area() float64{
	return r.width*r.height
}
func (r rect) perim() float64{
	return 2*r.width+2*r.height
}

func (c circle) area() float64{
	return math.Pi*c.radius*c.radius
}

func (c circle) perim() float64{
	return 2*math.Pi*c.radius
}

func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// Error
func f(arg int)(int,error){
	if arg==42{
		return -1,errors.New("can't work with 42")
	}
	return arg+3,nil
}

var ErrorOutOfTea=fmt.Errorf("no more tea available")
var ErrPower=fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg==2 {
		return ErrorOutOfTea
	} else if arg==4 {
		return fmt.Errorf("making tea:%w",ErrPower)
	}
	return nil
}

// Goroutine
func fn(from string){
	for i:=0;i<3;i++{
		fmt.Println(from,":",i)
	}
}

func main(){
	// 1 -> Closure
	// nextInt:=intSeq()
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())

	//2 -> Recursion
	// var num int
	// fmt.Println("Enter your number:")
	// fmt.Scanln(&num)  // Taking input 

	// fmt.Println("factorial of num is:",fact(num))


	// // anonymous function
	// var fib func(n int) int

	// fib=func(n int) int{
	// 	if n<2 {
	// 		return n
	// 	}
	// 	return fib(n-1)+fib(n-2)
	// }
	// fmt.Println("Fibonacci is:",fib(7))


	// 3 -> Range over built-in type
	// nums:=[]int{2,3,4,5}
	// sum:=0
	// for _, num :=range nums{
	// 	sum+=num
	// }
	// fmt.Printf("sum is:%d\n",sum)

	// kvs:=map[string]string{"a":"apple","b":"banana"}
	// for k,v:=range kvs{
	// 	fmt.Printf("%s -> %s\n",k,v)
	// }

	// i:=1
	// fmt.Println("initial:",i)
	// zeroval(i)
	// fmt.Println("zeroval:",i)

	// zeroptr(&i)
	// fmt.Println("zeroptr:",i)

	// fmt.Println("pointer:",&i)


	// 4 -> String and ruins
	// const s = "สวัสดี"
	// fmt.Println("Len:",len(s))

	// for i:=0;i<len(s);i++{
	// 	fmt.Println("%x",s[i])
	// }
	// fmt.Println()
	// fmt.Println("Rune count:",utf8.RuneCountInString(s))


	// for idx,runeValue:=range s{
	// 	fmt.Printf("%#U starts at %d\n",runeValue,idx)  // Here %#U is used for format rune for displaying both rune and it's character
	// }

	// 5-> struct
	// fmt.Println(person{"Bob",20})
	// fmt.Println(person{name:"Alice",age:30})
	// fmt.Println(person{name:"Fred"})
	// fmt.Println(newPerson("John"))


	// s:=person{name:"Sean",age:43}
	// fmt.Println(s.name)

	

	// 6 -> Methods
	// r:=rect{width:8,height:12}
	// fmt.Println("area",r.area())
	// fmt.Println("perimeter",r.perim())


	// rp:=&r
	// fmt.Println("area:",rp.area())
	// fmt.Println("perimeter:",rp.perim())



	// 7 -> Interfaces
	// r:=rect{width:3,height:4}
	// c:=circle{radius:5}

	// measure(r)
	// measure(c)


	// 8 -> Goroutines
	// fn("direct")
	// go fn("goroutine")

	// time.Sleep(time.Second)
	// fmt.Println("done")


	// // 9 => Channel
	// messages:=make(chan string)
	// go func(){
	// 	messages<-"ping"
	// }()

	// fmt.Println(messages)
	// msg := <-messages
    // fmt.Println(msg)

	// 10 => Timer
	timer1:=time.NewTimer(2*time.Second)
	fmt.Println(timer1)



	


	

	









	

	



}