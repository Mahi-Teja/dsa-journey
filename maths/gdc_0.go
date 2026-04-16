package maths

import "fmt"

func Gcd(){
	n := 36
	m := 12
	a := [] int {} 
	b := [] int {} 
	for i := 1; i*i <= n; i++ {
		if(n%i==0){
			a = append(a, i)
			if(i!=n/i){
				a= append(a, n/i)

			}
		}
	}
	for i := 1; i*i <= m; i++ {
		if(m%i==0){
			b = append(b, i)
			if(i!=n/i){
				
				b= append(b, m/i)

			}

		}
	}
	common := []int{}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if(a[i]==b[j]){
				common = append(common, a[i])
				// fmt.Println(i,j)
			}
		}
	} 
	gcd :=  0
	for i := 0; i < len(common); i++ {
		if(gcd<common[i]){
			gcd = common[i]
		}
	}
	fmt.Println(gcd)
}