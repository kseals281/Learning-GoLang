package main

import "fmt"

func main() {
  numbers := [10]int{2, 1, 5, 2, 3, 7, 6, 8, 4, 0}
  fmt.Println("Original list: ", numbers)
  for i := 0; i < len(numbers); i++ {
    for j := 0; j < len(numbers) - i - 1; j++ {
      if numbers[j] > numbers[j+1] {
        temp := numbers[j]
        numbers[j] = numbers[j+1]
        numbers[j+1] = temp
      }
    }
  }
  fmt.Println("Sorted list: ", numbers)
}