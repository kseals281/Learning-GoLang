package main

import "fmt"
import "math"

func getChange(amount float64) (lincoln, washington, quarters, dimes, nickels, pennies int) {
  for amount >= .01 { //One one currency of each type per loop
    if int(amount / 5) > 0 {
      lincoln++
      amount -= 5
    }
    if int(math.Mod(amount, 5) / 1) > 0 {
      washington++
      amount -= 1
    }
    if int(math.Mod(math.Mod(amount, 5), 1) / .25) > 0 {
      quarters++
      amount -= .25
    }
    if int(math.Mod(math.Mod(math.Mod(amount, 5), 1), .25) / .1) > 0{
      dimes++
      amount -= .1
    }
    if int(math.Mod(math.Mod(math.Mod(math.Mod(amount, 5), 1), .25), .1) / .05) > 0 {
      nickels++
      amount -= .05
    }
    if int(math.Mod(math.Mod(math.Mod(math.Mod(math.Mod(amount, 5), 1), .25), .1), .05) / .01) > 0 {
      pennies++
      amount -= .01
    }
  }
  return
}

func main() {
  var amount float64 = 3.69
  lincoln, washington, quarters, dimes, nickels, pennies := getChange(amount)
  fmt.Printf("$%.2f can be separated into %d fives, %d singles, %d quarters, %d dimes, ",
             amount, lincoln, washington, quarters, dimes)
  fmt.Printf("%d nickels, and %d pennies\n", nickels, pennies)
}