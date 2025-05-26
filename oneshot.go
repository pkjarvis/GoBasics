package main

import (
	"fmt"
	"strings"
	// "unicode/utf8"
	// "errors"
	// "math/rand"
	"time"
	"sync"
)

var wg=sync.WaitGroup{}
var dbData=[]string{"id1","id2","id3","id4","id5"}
var results=[]string{}
var m=sync.RWMutex{}

func dbCall(i int){
	var delay float32=2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is:",dbData[i])

	m.Lock()
	results=append(dbData,dbData[i])
	m.Unlock()

	wg.Done()
}


// func printMe(printVal string){
// 	fmt.Println(printVal)
// }


// func intDivision(numerator int,denominator int)(int ,int){

// 	var result int=numerator/denominator
// // 	var remainder int=numerator%denominator
// // 	return result,remainder
// // }

// func intDivision1(){
// 	var numerator int=10
// 	var denominator int =20
// 	var result,remainder= intDivision(numerator,denominator)
// 	switch{
// 	case err!=nil:
// 		fmt.Printf(err.Error())
	
//     case remainder==0:
// 	    fmt.Printf("The result of the integer division is%v",result)
// 	default:
// 		fmt.Printf("The result of the integer division is%v with remainder %v",result,remainder)
//     }
// }


func main(){
	// var intNum int =323
	// intNum++;
	// fmt.Println(intNum)

	// var floatnum float64=1234
	// fmt.Println(floatnum)

	// var floatnum32 float32=23
	// fmt.Println(floatnum32)

	// var result float32=floatnum32+ float32(intNum)
	// fmt.Println(result)

	// var intNum1 int=3
	// var intNum2 int=4
	// fmt.Println(intNum1/intNum2)


	// var mystring string = "Hello"+" "+"World"
	// fmt.Println(mystring)

	// fmt.Println(len(mystring))  // It basically stores character 


	// fmt.Println(utf8.RuneCountInString("Y"))

	// var myRune rune='a'
	// fmt.Println(myRune)

	// var myBoolean bool=false
	// fmt.Println(myBoolean)


	// // default value for int,uint is 0 , for run it is null , for boolean it is false
	// myVar:=3
	// fmt.Println(myVar)

	// var1,var2:=1,2
	// fmt.Println(var1,var2)


	// // printMe("hello from function")
	// // a,b:=intDivision(8,4)
	// // println(a,b)

	// intDivision1()



	arr:=[...]int64{1,2,3}
	fmt.Println(arr)


	var Mymap map[string]uint8=make(map[string]uint8)
	fmt.Println(Mymap)

	var myMap2 = map[string]uint8{"Adam":23,"Sarah":22}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Sarah"])
	

	var age,ok=myMap2["Jason"]
	if ok{
		fmt.Println("The age is %v",age)
	}else{
		fmt.Println("Invalid name")
	}

	for name,age:=range myMap2{
		fmt.Printf("Name:%v , Age:%v \n",name,age)
	}


	var myRune='a'
	fmt.Printf("myRune=%v",myRune)

	var strSlice=[]string{"s","u","b","t","c","r","i"}
	var b strings.Builder 
	for i:= range strSlice{
		b.WriteString(strSlice[i])
	}

	var catstr=b.String()
	fmt.Printf("%v\n",catstr)












	
	  





}