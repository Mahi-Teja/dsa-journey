package reccursion

import "fmt"


func PrintNumber(i int,n int){
	if(i>n) {
		return
	}

	fmt.Println(i)
	PrintNumber(i+1,n)
}
func PrintNumberDesc(i int,n int){
	if(i<1) {
		return
	}

	fmt.Println(i)
	PrintNumberDesc(i-1,n)
}