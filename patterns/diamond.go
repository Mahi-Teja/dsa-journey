package patterns

import "fmt"

func DiamondPattern(){
	n:= 5
	for row := 0; row < n; row++ {
		
		for space := 0; space < n-row-1; space++ {
			fmt.Print(` `)
		}
		for star := 0; star < 2*row+1; star++ {
			fmt.Printf(`%v`,row)
			
		}
		for space := 0; space < n-row-1; space++ {
			fmt.Print(` `)
			
		} 

		fmt.Println()
	}
	for row := n-1; row>=0; row-- {
		
		
		for space := n-row-1; space >0; space-- {
			fmt.Print(` `)
		}
		for star := 2*row+1; star >0; star-- {
			fmt.Printf(`%v`,row)
			
		}
		for space := n-row-1; space >0; space-- {
			fmt.Print(` `)
		}

		fmt.Println()
	}
}
/*
4,1,4
3,3,3
2,5,2
1,7,1
0,9,0
*/