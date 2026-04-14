package patterns

import "fmt"

func HalfPyramidStar( ){
num := 4 
for i := 1; i <= num; i++ {
	for j := 1; j <=i; j++ {	
		fmt.Print(`*`)
	}
	fmt.Println()
}

}
