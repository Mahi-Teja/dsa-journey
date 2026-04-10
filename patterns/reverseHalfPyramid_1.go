package main

import "fmt"

func reverseHalfPyramid_1(){ 
	n:=5
	for i := n; i > 0; i-- {
		for j := i; j >0; j-- {
			fmt.Print(`*`)
		}
		fmt.Println()
	}
}