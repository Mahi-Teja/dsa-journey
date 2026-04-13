package patterns

import "fmt"

func Pyramid(){
	n:= 10
	for i := 1; i <=n; i++ {
		for space := 0; space <= n-i-1; space++ {
		   fmt.Print(`   `)
		}
		for j := 1; j <=2*i-1; j++ {
			fmt.Print(` * `)
			
		}
		for space := 0; space <= n-i-1; space++ {
		   fmt.Print(`  `)
		}
			 
		fmt.Println()
	}
}
