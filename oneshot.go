package main

import (
	"fmt"
	// "strings"
	// "unicode/utf8"
	// "errors"
	// "math/rand"
	// "time"
	// "sync"
)

// var wg=sync.WaitGroup{}

// var MAX_CHICKEN_PRICE float32=5
// var dbData=[]string{"id1","id2","id3","id4","id5"}
// var results=[]string{}
// var m=sync.Mutex{}
// var v=sync.RWMutex{}

// func dbCall(i int){
// 	// Simulate Db call delay
// 	var delay float32=2000
// 	time.Sleep(time.Duration(delay)*time.Millisecond)  // it pauses the latest go routine for provided time delay
// 	fmt.Println("The result from the database is:",dbData[i])
// 	save(dbData[i])
// 	log()

// 	m.Lock()
// 	results=append(results,dbData[i])
// 	m.Unlock()

// 	wg.Done()
// }



// func log(){
// 	v.RLock()
// 	fmt.Println("\nThe current results are:%v",results)
// 	v.RUnlock()
// }

// func save(result string){
// 	m.Lock()
// 	results=append(results,result)
// 	m.Unlock()
// }


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


// func count(){
// 	var res int
// 	for i:=0;i<100000000;i++{
// 		res+=1;
// 	}
// 	wg.Done()
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



	// arr:=[...]int64{1,2,3}
	// fmt.Println(arr)


	// var Mymap map[string]uint8=make(map[string]uint8)
	// fmt.Println(Mymap)

	// var myMap2 = map[string]uint8{"Adam":23,"Sarah":22}
	// fmt.Println(myMap2["Adam"])
	// fmt.Println(myMap2["Sarah"])
	

	// var age,ok=myMap2["Jason"]
	// if ok{
	// 	fmt.Println("The age is %v",age)
	// }else{
	// 	fmt.Println("Invalid name")
	// }

	// for name,age:=range myMap2{
	// 	fmt.Printf("Name:%v , Age:%v \n",name,age)
	// }


	// var myRune='a'
	// fmt.Printf("myRune=%v",myRune)

	// var strSlice=[]string{"s","u","b","t","c","r","i"}
	// var b strings.Builder 
	// for i:= range strSlice{
	// 	b.WriteString(strSlice[i])
	// }

	// var catstr=b.String()
	// fmt.Printf("%v\n",catstr)

	// fmt.Println("end")
	// t0:=time.Now()
	// for i:=0;i<1000;i++{
	// 	wg.Add(1)
	// 	go count()
	// }
	// wg.Wait()
	// fmt.Println("\nTotal execution time:%v",time.Since(t0))


	// var c=make(chan int)
	// c<-1
	// var i=<-c   // this blocks the code execution and hence gives deadlock error
	// fmt.Println(i)

	// go process(c)
	// for i:=range c{
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second*1)
	// }

	// var chickenChannel=make(chan string)
	// var websites=[]string{"walmart.com","costco.com","wholefoods.com"}
	// for i:=range websites{
	// 	go checkChickenPrices(websites[i],chickenChannel)
	// }
	// sendMessage(chickenChannel)


	var intSlice=[]int{1,2,3}
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice=[]float32{1,2,3}
	fmt.Println(sumSlice[float32](float32Slice))



	
	
}

// func process(c chan int){
// 	defer close(c)
// 	for i:=0;i<5;i++{
// 		c<-i
// 	}
// 	// close(c)
// 	fmt.Println("existing process")
// }


// func checkChickenPrices(website string,chickenChannel chan string){
// 	for{
// 		time.Sleep(time.Second*1)
// 		var checkChickenPrice=rand.Float32()*20
// 		if checkChickenPrice<=MAX_CHICKEN_PRICE{
// 			chickenChannel<-website 
// 			break
// 		}
// 	}
// }

// func sendMessage(chickenChannel chan string){
// 	fmt.Printf("\nFound a deal on chicken at %s",<-chickenChannel)
// }

func sumSlice[T int|float32|float64](slice []T)T{
	var sum T
	for _,v:=range slice{
		sum+=v
	}
	return sum

}