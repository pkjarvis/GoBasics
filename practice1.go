package main

// Package runtime contains operations that interact with Go's runtime system, such as functions to control goroutines.
// It also includes the low-level type information used by the reflect package;
import (
	"fmt"
	"math"
	"runtime"
	"time"
	"strings"

)

type Vertex struct{
	X int 
	Y int
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	// fmt.Sprint is primarily used for building strings by joining different data types together
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow1(x, n, lim float64) float64 {
	// variables declared inside v are also available in else
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)


func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s)  // both len,cap are for declaring len,capacity of array
}


var pow2=[]int{1,2,4,8,16,32,64,128}



func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// closure ->  A closure is a function value that references variables from outside its body. 
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}



type VertexV struct{
	X,Y float64
}

func (v VertexV) Abs() float64{
	return Math.Sqrt(v.X*v.X+v.Y*v.Y)
}

type MyFloat float64
func (f MyFloat) Abs() float64{
	if f<0 {
		return float64(-f)
	}
	return float64(f)
}



// Define a struct type
type Rectangle struct {
	Width  float64
	Height float64
}

// Area is a method associated with the Rectangle type.
// It has a value receiver `r Rectangle`.
// This means the method operates on a COPY of the Rectangle value.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter is another method for Rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}



import "fmt"
import "math"

// Define an interface called 'Shape'
// Any type that has both an Area() and a Perimeter() method (with matching signatures)
// automatically satisfies the Shape interface.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle type (from previous example)
type Rectangle struct {
	Width  float64
	Height float64
}

// Area method for Rectangle (value receiver)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter method for Rectangle (value receiver)
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle type
type Circle struct {
	Radius float64
}

// Area method for Circle (value receiver)
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter method for Circle (value receiver - for a circle, this is circumference)
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// PrintShapeInfo is a function that takes an argument of type 'Shape'.
// It can accept any type that implements the Shape interface.
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}



func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum is", sum)

	num := 0
	for num < 3 {
		num += num
		num += 1
	}
	fmt.Println("num is:", num)

	// for is go's while
	x := 1
	for x < 1000 {
		x += x
	}
	fmt.Println("x :", x)

	// infinite loop is written like this -
	// for {
	// }

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	fmt.Println(pow1(3, 2, 10), pow1(3, 3, 20))

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	defer fmt.Println("world") // Used to give delay , it delays the execution of function until surrounding function returns

	fmt.Println("hello")

	// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	fmt.Println("Hi")
	// Pointers -> problems not output in console
	i, j := 42, 27

	p := &i
	fmt.Println(*p)
	fmt.Println(p)

	*p = 21
	fmt.Println(i)

	p = &j
	*p=*p/10
	fmt.Println(j)

	fmt.Println("checking")

	fmt.Println("Vertex is:",Vertex{1,2})

	v:=Vertex{3,4}
	v.X=4
	fmt.Println("vertex x is",v.X)

	q:=&v 
	q.X=1e9
	fmt.Println(v)


	fmt.Println("struct literals",v1,p,v2,v3)



	// Array
	var a[2] string
	a[0]="Hello"
	a[1]="World"

	fmt.Println(a[0],a[1])
	fmt.Println(a)

	var primes=[6]int{2,3,4,5,6,8}  // or primes:=[6]int{2,3,4,5,6,7}
	fmt.Println(primes)

	var s[]int=primes[1:4]
	fmt.Println(s)

	names:=[4]string{
		"John","Paul","George","Ringo",
	}

	fmt.Println(names)

	y:=names[0:2]
	z:=names[1:3]
	fmt.Println(y,z)

	y[0]="abc"
	fmt.Println(names)



	l:=[]int{2,3,4,5,8}
	fmt.Println("Arrays",l)

	m:=[]bool{true,false,true,true,false}
	fmt.Println(m)


	n:=[]struct{
		i int
		b bool
	}{
		{2,true},
		{3,false},
		{5,true},
		{7,false},
		{8,true},
		{9,true},
	}


	fmt.Println(n)


	px:=[]int{2,3,4,7,8,11,13}
	var u=px[1:5]
	fmt.Println(u)

	var tx=px[:2]   // by default it takes values from 0 for both lower as well as upper bound
	fmt.Println(tx)



	printSlice(px)
	printSlice(tx)



	fmt.Println(px,len(px),cap(px))

	if px==nil{
		fmt.Println("nil!")
	}


	ax:=make([]int,6)
	printSlice(ax)

	bx:=make([]int,0,5)
	printSlice(bx)


	// Create a tic-toe board
	board:=[][]string{
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}

	// The players take the turns.
	board[0][0]="X"
	board[2][2]="O"
	board[1][2]="X"
	board[1][0]="O"
	board[0][2]="X"

	fmt.Println("board len:",len(board))

	for i:=0;i<len(board);i++ {
		fmt.Printf("%s\n",strings.Join(board[i]," "))
	}


	var vx []int
	vx=append(vx,0)  // append work on nil slices, it add element to end of array element
	printSlice(vx)
	vx=append(vx,1)
	printSlice(vx)

	vx=append(vx,2,3,4)
	printSlice(vx)


	for n:=range 6{
		if n%2==0{
			continue;
		}
		fmt.Println(n)
	}


	for i,v:= range pow2{
		fmt.Printf("2**%d=%d\n",i,v)
	}


	pow3:=make([]int,10)
	for i:=range pow3 {
		pow3[i]= 1<< uint(i) // ===2**i
	}

	for _,value:=range pow3{
		fmt.Printf("%d\n",value)
	}


	// Map

	mp:=make(map[string]int)

	mp["k1"]=7
	mp["k2"]=13

	fmt.Println("maps:",mp)

	v2:=mp["k1"]
	fmt.Println("map k1 key value:",v2)

	v3:=mp["k3"]
	fmt.Println("v3:",v3)

	delete(mp,"k2")
	fmt.Println("map:",mp)

	clear(mp)
	fmt.Println(mp)

	nx:=map[string]int{"foo":1,"bar":2}
	fmt.Println("map",nx)

	ny:=map[string]int{"foo":1,"bar":2}
	fmt.Println("map",ny)


	hypot:=func(x,y float64) float64{
		return math.Sqrt(x*x+y*y)
	}

	fmt.Println("Sqrt of num x,y is:",hypot(5,12))
	fmt.Println("Compute value is:",compute(hypot))
	fmt.Println("Power is:",compute(math.Pow))

	pos,neg:=adder(),adder()
	for i:=0;i<10;i++ {
		fmt.Println(pos(i),neg(-2*i))

	}


	//  Method syntax
	func (receiverName ReceiverType) MethodName(parameters) (returnValues) {
    // Method body
    }

	// Interfaces in go
	// An interface defines a set of method signatures. It specifies what a type can do, not how it does it
	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}

	// Both Rectangle and Circle satisfy the Shape interface
	// because they both have Area() and Perimeter() methods.
	PrintShapeInfo(rect) // Output: Area: 50.00, Perimeter: 30.00
	PrintShapeInfo(circ) // Output: Area: 153.94, Perimeter: 43.98

	// You can also declare interface variables directly
	var s1 Shape = rect
	var s2 Shape = circ
	fmt.Println("Using interface variables:")
	PrintShapeInfo(s1)
	PrintShapeInfo(s2)

	

	

	

	



	





	

	








	



	

















	










}
