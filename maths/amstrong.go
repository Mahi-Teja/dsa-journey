package maths

import "fmt"

func Amstrong(){
	n:=1634
	org := n
	amst :=0
	for n>0{
		digit := n%10
		cb := digit * digit * digit
		fmt.Println( cb)
		amst = amst + ( cb)
		
		// fmt.Println( amst)
		n = n/10
	}
	// fmt.Print( amst)
	fmt.Print(org == amst)
}