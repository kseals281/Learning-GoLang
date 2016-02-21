package main

import "fmt"
import "sort"

func main() {
  numbers := [5]int{3, 6, 2, 1, 5}
  for i := 0; i < len(numbers); i++ {
    for j := 0 + i; j < len(numbers) - 1; j++ {
      if numbers[j] > numbers[j+1] {
        sort.Swap(numbers[j], numbers[j + 1])
      }
    }
  }
  fmt.Println(numbers)
}