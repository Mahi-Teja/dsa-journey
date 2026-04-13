package patterns

import "fmt"

func HalfPyramidStarNumber_2(){
	n:=5
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
}