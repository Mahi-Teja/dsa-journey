package reccursion

import "fmt"


func BTNumbers(i int,n int){

	if(n<1){
		return
	}
	BTNumbers(i,n-1)
	fmt.Println(n)
}
func BTNumbersRev(i int,n int){

	if(i>n){
		return
	}

	BTNumbersRev(i+1,n)
	fmt.Println(i)
}
