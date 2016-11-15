package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sherlockandthebeast(N int) int {
	var fives, threes int
	for (N % 3) != 0 {
		N -= 5
		fives++
	}
	if N < 0 {
		return -1
	}
	threes = N / 3
	c := []string{
		strings.Repeat("5", fives),
		strings.Repeat("3", threes),
	}
	result := strings.Join(c, "")
	i, _ := strconv.Atoi(result)
	return i
}

func main() {
	var T, N int
	fmt.Scanf("%d", &T)
	for t := 0; t < T; t++ {
		fmt.Scanf("%d", &N)
		fmt.Println(sherlockandthebeast(N))
	}
}
