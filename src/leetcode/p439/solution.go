package main

import "fmt"

func main() {
	//fmt.Println(parseTernary("T?2:3"))
	fmt.Println(parseTernary("F?1:T?4:5"))
	fmt.Println(parseTernary("T?T?F:5:3"))
}

func parseTernary(expression string) string {
	if len(expression) == 1 {
		return expression
	}
	condition, left, right := split(expression)

	condition = parseTernary(condition)
	if condition == "T" {
		return parseTernary(left)
	}
	return parseTernary(right)
}

func split(expr string) (string, string, string) {
	condition := ""
	for i := 0; i < len(expr); i++ {
		if expr[i] == '?' {
			condition = expr[:i]
			expr = expr[i+1:]
			break
		}
	}

	level := 0
	left, right := "", ""
	for i := 0; i < len(expr); i++ {
		if level == 0 && expr[i] == ':' {
			left, right = expr[:i], expr[i+1:]
			break
		}

		if expr[i] == '?' {
			level++
		} else if expr[i] == ':' {
			level--
		}
	}

	return condition, left, right
}
