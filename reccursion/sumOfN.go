package reccursion

import "fmt"

var Sum = 0
func SumOfN(i int,n int, ){

	 fmt.Println(sumIt(n))

}

func sumIt(n int) int{
if(n==0){
	return 0
}	
 return n + sumIt(n-1)
}