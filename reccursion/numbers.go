package reccursion

import "fmt"


func PrintNumber(i int,n int){
	if(i>n) {
		return
	}

	fmt.Println(i)
	PrintNumber(i+1,n)
}