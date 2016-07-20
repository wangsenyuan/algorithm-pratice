package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(isNumber(" 0.1 "))
	fmt.Println(isNumber("abc"))
}

const (
	InValid = iota
	Space
	Sign
	Num
	Dot
	E
)

var st = [][]int{
	[]int{-1, 0, 3, 1, 2, -1}, // next state for
	// state 0
	[]int{-1, 8, -1, 1, 4, 5},    // next states for state 1
	[]int{-1, -1, -1, 4, -1, -1}, // next states for state 2
	[]int{-1, -1, -1, 1, 2, -1},  // next states for state 3
	[]int{-1, 8, -1, 4, -1, 5},   // next states for state 4
	[]int{-1, -1, 6, 7, -1, -1},  // next states for state 5
	[]int{-1, -1, -1, 7, -1, -1}, // next states for state 6
	[]int{-1, 8, -1, 7, -1, -1},  // next states for state 7
	[]int{-1, 8, -1, -1, -1, -1}, // next states for state 8
}

func isNumber(s string) bool {
	state := 0

	for _, c := range s {
		x := InValid
		if unicode.IsSpace(c) {
			x = Space
		} else if c == '+' || c == '-' {
			x = Sign
		} else if unicode.IsDigit(c) {
			x = Num
		} else if c == '.' {
			x = Dot
		} else if c == 'e' {
			x = E
		}
		if x == InValid {
			return false
		}
		state = st[state][x]
		if state < 0 {
			return false
		}
	}

	return state == 1 || state == 4 || state == 7 || state == 8
}
