package main

import "fmt"

func main() {
  x := 3
  a, b := 1, 1
  c := 0
  if x == 1 || x == 2 {
    c = 1
  }
  for i := 0; i < x - 2; i++ {
    c = a + b
    a = b
    b = c
  }
  fmt.Printf("The %d Fibonacci number is %d\n", x, c)
}