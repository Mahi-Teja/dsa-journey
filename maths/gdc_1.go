package maths

import "fmt"

// Euclidian Algorithm
func Gcd_EA(){
	a:=12
	b:=36
	res:=0

	for a>0 && b>0 {
		if(a>b){
	a=a-b
}else{
	b=b-a
}
	}
	if(a==0){
		res=b
	}else {
		res = a
	}
	fmt.Println(res)
}


func Gcd_EA2(){
	a:=14
	b:=144
	res:=0

	for a>0 && b>0 {
		if(a>b){
	a=a%b
	}else{
		b=b%a
	} 
	}
	if(a==0){
		res=b
	}else {
		res = a
	}
	fmt.Println(res)
}

