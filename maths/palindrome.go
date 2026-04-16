package maths

import "fmt"

func Palindrome(){
	n:=1
	orginal := n
	rev := 0
	for n>0{
		digit:=n%10
		rev = rev*10 + digit
		n=n/10
	} 
	fmt.Print(rev==orginal)

}