package main

import "fmt"

func halfPyramidStar( ){
num := 4 
for i := 1; i <= num; i++ {
	for j := 1; j <=i; j++ {	
		fmt.Print(j)
	}
	fmt.Println()
}

}
