// Package is a way to group functions and it's made up of all the files in the same directory
package main

// importing fmt package  which contains  functions for formatting text, including printing to the console
import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}



func main(){
	fmt.Println("hello world")
	fmt.Println(m)
}



