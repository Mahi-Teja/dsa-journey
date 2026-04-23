package reccursion

import (
	"fmt"
	"strings"
)
 

func IsPalindrome(a string ) {
 

var arr = strings.Split(strings.ToUpper(a), "") // case in-sensitive

 fmt.Println(check(arr,0))
}

func check(arr []string, s int) bool  {
	if(s>len(arr)/2) {return true }  
	return   arr[s]==arr[len(arr)-s-1] && check(arr,s+1)
}