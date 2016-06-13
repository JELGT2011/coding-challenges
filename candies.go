package main

import "fmt"

var candies []int
var scores []int

func candiesperchild(index int) int {

  if candies[index] != 0 {
    return candies[index]
  }

  candidate := 0
  candy := 1
  if index + 1 < len(scores) && scores[index] > scores[index + 1] {
    candidate = candiesperchild(index + 1) + 1
    if candidate > candy { candy = candidate }
  }
  if index > 0 && scores[index] > scores[index - 1] {
    candidate = candiesperchild(index - 1) + 1
    if candidate > candy { candy = candidate }
  }
  candies[index] = candy
  return candies[index]
}

func main() {
  var T int
  fmt.Scanf("%d", &T)
  candies = make([]int, T)
  scores = make([]int, T)
  for t := 0; t < T; t++ {
    fmt.Scanf("%d", &scores[t])
  }

  for t := 0; t < T; t++ {
    candiesperchild(t)
  }

  sum := 0
  for t := 0; t < T; t++ {
    sum += candies[t];
  }
  fmt.Println(sum)
}
