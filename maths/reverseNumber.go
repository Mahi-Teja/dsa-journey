package maths

import "fmt"

func ReverseNumber(){
	n:=9702
	reverse := 0

	for n>0 {
		digit := (n%10)
		reverse = (reverse *10) + digit 
		n=n/10
	}

	fmt.Print(reverse)
}

//  Tc : O(log10 n)
