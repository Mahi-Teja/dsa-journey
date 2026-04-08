package main

import "fmt"

func reverseHalfPyramid(){ 
	n:=5
	for i := n; i > 0; i-- {
		for j := i; j >0; j-- {
			fmt.Print(i)
		}
		fmt.Println()
	}
}