package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(myAtoi("aaaa"))
}

const (
	INT_MAX int64 = 2147483647
	INT_MIN int64 = -2147483648
)

func myAtoi(str string) int {
	var y int64

	headSpace := true
	var sign int64 = 1
	signRead := false
	for _, r := range str {
		if y > INT_MAX {
			break
		}

		if headSpace && unicode.IsSpace(r) {
			continue
		}

		headSpace = false

		if !signRead && isPosSign(r) {
			sign = 1
			signRead = true
			continue
		}

		if !signRead && isNegSign(r) {
			sign = -1
			signRead = true
			continue
		}

		if unicode.IsDigit(r) {
			y = y*10 + int64(r-'0')
			continue
		}

		break
	}
	y = sign * y
	if y > INT_MAX {
		y = INT_MAX
	} else if y < INT_MIN {
		y = INT_MIN
	}

	return int(y)
}

var posSign = '+'
var negSign = '-'

func isPosSign(r rune) bool {
	return r == posSign
}

func isNegSign(r rune) bool {
	return r == negSign
}
