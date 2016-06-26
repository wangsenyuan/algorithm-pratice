package main

import "fmt"

func isPalindrome(x int) bool {
  if x < 0 {
    return false
  }

  y := 0
  _x := x
  for _x > 0 {
    z := _x % 10
    _x = _x / 10
    y = y * 10 + z
  }

  return x == y
}

func main() {
  fmt.Println(isPalindrome(1))
}
