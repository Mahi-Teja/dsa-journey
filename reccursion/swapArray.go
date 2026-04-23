package reccursion

import "fmt"
var arr = []int{1,5,6,22,56,43,99}
var length = len(arr)

func SwapArr(){
 
	fmt.Println(arr)
	Swap2(0)
	fmt.Println(arr)
}

// with 2 variables
func Swap(a int, b int)  { 
	if(a>b || a==b){
		return
	}
	first:=arr[a]
	arr[a]=arr[b]
	arr[b] = first
	
	Swap(a+1,b-1)
}
// with 1 variable
func Swap2(n int){
	var limit = len(arr)/2 
 
if(n>limit){
	return
}
	first:= arr[n]
	arr[n] = arr[length-n-1]
	arr[length-n-1] = first

	Swap2(n+1)

}
