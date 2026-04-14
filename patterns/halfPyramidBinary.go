package patterns

import "fmt"


func BinaryHalfPyramid(){
	n:=5
	start := 1
	for row := 0; row < n; row++ {
		if(row % 2==0){
			start = 1
			}else{
				start = 0
			}
		for col := 0; col <= row; col++ {
			fmt.Print(start)
			start = 1-start
		}
		fmt.Println()
	}
}