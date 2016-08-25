package main

import (
	"fmt"
	"strconv"
)

func main() {
	//tokens := []string{"2", "1", "+", "3", "*"}
	tokens := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(tokens))
}

func evalRPN(tokens []string) int {
	stack := make([]int, len(tokens))
	p := 0

	for _, token := range tokens {
		switch token {
		case "+":
			x := stack[p-1]
			p--
			y := stack[p-1]
			p--
			stack[p] = x + y
			p++
		case "-":
			x := stack[p-1]
			p--
			y := stack[p-1]
			p--
			stack[p] = y - x
			p++
		case "*":
			x := stack[p-1]
			p--
			y := stack[p-1]
			p--
			stack[p] = x * y
			p++
		case "/":
			x := stack[p-1]
			p--
			y := stack[p-1]
			p--
			stack[p] = y / x
			p++
		default:
			stack[p], _ = strconv.Atoi(token)
			p++
		}
	}

	return stack[0]
}
