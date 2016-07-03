package main

import "fmt"

func max(a int, b int) int {
  if a > b {
    return a
  }
  return b
}

func sum(A []int) int {
  var s int
  for i := 0; i < len(A); i++ {
    s += A[i]
  }
  return s
}

func arraysplitting(A []int, moves int) int {
  if len(A) < 2 {
    return moves
  }
  for i := 1; i < len(A); i++ {
    if sum(A[:i]) == sum(A[i:]) {
      return max(arraysplitting(A[:i], moves + 1), arraysplitting(A[i:], moves + 1))
    }
  }
  return moves
}

func main() {
  var T, N int
  var A []int
  fmt.Scanf("%d", &T)
  for t := 0; t < T; t++ {
    fmt.Scanf("%d", &N)
    A = make([]int, N)
    for n := 0; n < N; n++ {
      fmt.Scanf("%d", &A[n])
    }
    fmt.Println(arraysplitting(A, 0))
  }
}
