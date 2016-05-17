package main

import "fmt"

var cache map[int]int
var weights map[int]bool

func unboundedknapsack(k int) int {
	if val, ok := cache[k]; ok {
		return val
	}

	var s int
	best := k
	for weight := range weights {
		if k-weight >= 0 {
			s = knapsack(k - weight)
			if s < best {
				best = s
			}
		}
	}
	cache[k] = best
	return cache[k]
}

func main() {
	var T, N, K int
	var weight int
	fmt.Scanf("%d", &T)
	for t := 0; t < T; t++ {
		fmt.Scanf("%d %d", &N, &K)
		cache = make(map[int]int)
		weights = make(map[int]bool)
		for n := 0; n < N; n++ {
			fmt.Scanf("%d", &weight)
			weights[weight] = true
		}
		fmt.Println(K - unboundedknapsack(K))
	}
}
