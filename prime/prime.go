package main

import (
  "fmt"
  "math"
)

func main() {
  var x float64 = 9.0
  prime := true
  var out_string string = ""
  for i := 0.0; i < x; i++ {
    if math.Mod(x, i) == 0 {
      prime = false
    }
  }
  var y int = int(x)
  if prime == false {
    out_string = "not "
  }
  fmt.Printf("%d is %sa prime number\n", y, out_string)
}