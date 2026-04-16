package maths

import "fmt"

func PrimeNumber(){
	n:=8
	count := 0
	if(n<2) {
		 fmt.Print(`Number should be greater than 1`)
		 return
	}
	for i := 2; i*i < n; i++ { 
		if(n%i==0){
			count++
		}  
	} 
	if(count>0){

		fmt.Printf(`%v is a not a Prime`,n,)
		}else {
			
			fmt.Printf(`%v is a Prime`,n,)
	}

}