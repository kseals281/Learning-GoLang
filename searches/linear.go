package main

import "fmt"

func main() {
  var num int = 43
  var present bool = false
  numbers := [10]int{2, 1, 5, 2, 3, 7, 6, 8, 4, 0}
  for i := 0; i < len(numbers); i++ {
    if num == numbers[i] {
      present = true
    }
  }
  if present {
    fmt.Printf("%d is in the list\n", num)
  } else {
    fmt.Printf("%d is not in the list\n", num)
  }
}