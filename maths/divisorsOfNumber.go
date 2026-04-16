package maths

import (
	"fmt"
)

func DivisorsOfN(){
	n:=36

	for i:=1;i*i<=n ;i++{
		if(n%i==0){
			fmt.Println(i)
			if(n/i != i){
				
				fmt.Println(n/i)
			}
		}
	}
}