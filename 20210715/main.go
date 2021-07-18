package main

import (
	"fmt"
	"./util"
)


func main() {
	printHeader()

	A := util.NewSet("1", "2")
	B := util.NewSet("3", "4")
	C := util.Union(A, B)
	fmt.Println(C)

	// util.Intersection(A, B)
}


func printHeader(){
	fmt.Println("#######################")
	fmt.Println("### Set Operations ####")
	fmt.Println("#######################")
}