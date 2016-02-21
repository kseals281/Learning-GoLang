package main

import (
  "fmt"
  "math"
)

func main() {
  var x int = 28
  prime := true
  var out_string string = ""
  for i := 2; i*i < x; i++ {
    if math.Mod(float64(x), float64(i*i)) == 0 {
      prime = false
    }
  }
  if prime == false {
    out_string = "not "
  }
  fmt.Printf("%d is %sa prime number\n", int(x), out_string)
}