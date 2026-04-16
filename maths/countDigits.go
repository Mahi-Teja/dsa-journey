package maths

import "fmt"


func CountDigits(){
n:=78788787

count := 0
for n>0 { 
		count++ 
	n = n/10
}
fmt.Print(count)
}

// Tc : O(log10 n) 