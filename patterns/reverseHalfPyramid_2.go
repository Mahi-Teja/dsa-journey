package main

import "fmt"

func reverseHalfPyramid_2(){ 
	n:=5
	for i := n; i > 0; i-- {
		for j := i; j >0; j-- {
			fmt.Print(j)
		}
		fmt.Println()
	}
}