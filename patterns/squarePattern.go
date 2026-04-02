package main

import "fmt"

func main(){
	var n = 5
	for i :=1;i<=n;i++{
		for j := 1; j <= n; j++ {
			fmt.Printf(`*`) // use ` ` for printing string.
		}
		fmt.Println()
	}

}