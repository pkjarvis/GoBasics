package main

import (
	"fmt"
	"math/rand"
	"math"
	"time"
	"math/cmplx"
	
)

var (
	ToBe bool=false
	MaxInt uint64 =1<<64-1
	z   complex128=cmplx.Sqrt(-5+12i)
)

const s string ="constant"

var python,java bool
var ax,bx int =1,2


func add(x int,y int) int {
	return x+y
}

func swap(x, y string) (string, string) {
    return y, x
}

// Named return values
func split(sum int)(x,y int){
	x=sum*4/9
	y=sum-x
	return
}



func main(){
	fmt.Println("How u doing?");

	fmt.Println("3+2=",3+2,",",3+2,"Sum is",3+2);
	fmt.Println("7/3=",7/3);
	fmt.Println("7/3=",7.0/3.0);
	fmt.Println(true && false);
	fmt.Println(true || false);
	fmt.Println(!true);


	// Variables
	var a="initial"
	fmt.Println(a)

	var b,c int  = 1,2 // while declaraing varialbe we can assign types at the same time , if not mentioned it would be int automatically else for each we have to declare
	fmt.Println(b,c);

	var A=true
	fmt.Println(A)

	var e int 
	fmt.Println(e)

	// Shorthand version of declaring and initializing variable
	f:="apple"
	fmt.Println(f)

	const n=5e7
	const d=3e20/n;
	fmt.Println(d)

	fmt.Println(int64(d))  // print till complete division with decimal


	//  Maths 
	fmt.Println("Math function")

	fmt.Println(math.Sin(n))
	fmt.Println(math.Pi)
	fmt.Println(add(12,32))


	// For loop in go

	i:=1;
	for i<=3 {
		fmt.Println(i)
		i=i+1
	}

	for j:=0;j<3;j++{
		fmt.Println("Print",j)
	}

	// Range is only valid for iteratable like map,array,slice
	for i := range []int{2,3,4} {
		fmt.Println("range",i)
	}

	// for n := range 6 {
    //     if n%2 == 0 {
    //         continue
    //     }
    //     fmt.Println(n)
    // }

	fmt.Println("Square root of 9 is",math.Sqrt(9))

	// if/else
	if 75%2==0 {
		fmt.Println("7 is odd")
	}else{
		fmt.Println("7 is even")
	}


	C,D := swap("hello", "world")
	fmt.Println(C,D)

	fmt.Println("The time is", time.Now())

	rand.Seed(time.Now().UnixNano())    // This basically seed with current time such that rand would output random number between defined range
	fmt.Println("My favorite number is",rand.Intn(10))


	fmt.Println(split(17))

	// Whenever you define variable it always gets some value instead of garbage value
	var cx int
	fmt.Println(cx,python,java)

	


	if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

	//  Data Types -> bool,string,int,int 8,int 16,int 32,int 64,uint,uint8,uint32,uint64,uintptr,byte,rune,float32,float64,complex64,complex128


// 	%v	Default format (any value)
// %t	Boolean
// %#v	Go-syntax representation
// %T	Type of value


	fmt.Printf("Type: %T Value: %v\n",ToBe,ToBe)
	fmt.Printf("Type: %T Value: %v\n",MaxInt,MaxInt)
	fmt.Printf("Type: %T Value: %v\n",z,z)


	// %v	Default format	fmt.Printf("%v", val)
// %+v	Struct fields with names	fmt.Printf("%+v", struct)
// %#v	Go-syntax representation	fmt.Printf("%#v", val)
// %T	Type of the value	fmt.Printf("%T", val)
// %%	A literal percent sign	fmt.Printf("%%")

// Verb	Description	Example
// %d	Decimal	fmt.Printf("%d", 123)
// %b	Binary	fmt.Printf("%b", 5)
// %o	Octal	fmt.Printf("%o", 8)
// %x	Hex (lowercase)	fmt.Printf("%x", 255)
// %X	Hex (uppercase)	fmt.Printf("%X", 255)
// %c	Unicode character	fmt.Printf("%c", 65)
// %U	Unicode format	fmt.Printf("%U", 65)

// Verb	Description	Example
// %f	Decimal point	fmt.Printf("%f", 3.14)
// %e	Scientific notation (e)	fmt.Printf("%e", 3.14)
// %E	Scientific notation (E)	fmt.Printf("%E", 3.14)
// %g	Shortest of %e or %f	fmt.Printf("%g", 3.14)
// %G	Shortest of %E or %f	fmt.Printf("%G", 3.14)


// Verb	Description	Example
// %s	String	fmt.Printf("%s", "Go")
// %q	Quoted string	fmt.Printf("%q", "Go")
// %x	Hex (bytes of string)	fmt.Printf("%x", "Hi")
// %X	Uppercase hex	fmt.Printf("%X", "Hi")

var p,q int =3,4
var r float64=math.Sqrt(float64(p*p+q*q))

fmt.Println(p,q,r)




















}



