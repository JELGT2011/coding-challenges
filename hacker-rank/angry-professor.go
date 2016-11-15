package main

import "fmt"

func angryprofessor(A []int, K int) string {
	var count int
	for _, a := range A {
		if a <= 0 {
			count++
		}
	}
	if count < K {
		return "YES"
	}
	return "NO"
}

func main() {
	var T, N, K int
	var A []int
	fmt.Scanf("%d", &T)
	for t := 0; t < T; t++ {
		fmt.Scanf("%d", &N)
		fmt.Scanf("%d", &K)
		A = make([]int, N)
		for n := 0; n < N; n++ {
			fmt.Scanf("%d", &A[n])
		}
		fmt.Println(angryprofessor(A, K))
	}
}
