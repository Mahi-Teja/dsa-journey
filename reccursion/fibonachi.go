package reccursion

func Fibo (n int) int{

	if(n==0) { return 0}
	if(n==1 || n==2) { return 1}

	return  Fibo(n-1) + Fibo(n-2) 

}

// 0,1,1,2,3,5,8,13,21,34..
