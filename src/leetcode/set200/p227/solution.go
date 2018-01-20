package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calculate("161+6033/2/69-11*2-12*6+51*40"))
}

func calculate(s string) int {
	tokens := tokenize(s)
	nums := make([]int, len(tokens))
	var np int
	ops := make([]string, len(tokens))
	var op int

	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]
		if tok == "+" || tok == "-" || tok == "*" || tok == "/" {
			for op > 0 && prefer(ops[op-1], tok) {
				tmp := eval(ops[op-1], nums[np-2], nums[np-1])
				np -= 2
				nums[np] = tmp
				np++
				op--
			}
			ops[op] = tok
			op++
		} else {
			val, _ := strconv.Atoi(tok)
			nums[np] = val
			np++
		}
	}
	for op > 0 {
		tmp := eval(ops[op-1], nums[np-2], nums[np-1])
		np -= 2
		nums[np] = tmp
		np++
		op--
	}

	return nums[0]
}

func eval(op string, a int, b int) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	panic("wrong")
}

func tokenize(s string) []string {
	bs := []byte(s)
	var sz int
	for i := 0; i < len(bs); i++ {
		if bs[i] == ' ' {
			continue
		}
		bs[sz] = bs[i]
		sz++
	}

	bs = bs[:sz]

	res := make([]string, 0, 10)
	for i, j := 0, 0; i <= sz; i++ {
		if i == sz || bs[i] == '+' || bs[i] == '-' || bs[i] == '*' || bs[i] == '/' {
			res = append(res, string(bs[j:i]))
			if i < sz {
				res = append(res, string(bs[i]))
			}
			j = i + 1
		}
	}

	return res
}

func prefer(a, b string) bool {
	if a == "*" || a == "/" {
		return true
	}

	if b == "+" || b == "-" {
		return true
	}

	return false
}
