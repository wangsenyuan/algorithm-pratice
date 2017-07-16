package main

import (
	"strings"
	"fmt"
)

func solveEquation(equation string) string {
	problem := strings.Split(equation, "=")
	a, b := normalize(problem[0])
	c, d := normalize(problem[1])

	if a == c {
		if b == d {
			return "Infinite solutions"
		} else {
			return "No solution"
		}
	} else {
		ans := (d - b) / (a - c)
		return fmt.Sprintf("x=%d", ans)
	}
}

func normalize(str string) (int, int) {
	coefficient := 0
	constants := 0
	hasLeading := false
	num := 0
	symbol := 1
	for i := 0; i < len(str); i++ {
		if str[i] == 'x' {
			if !hasLeading && num == 0 {
				num = 1
			}
			coefficient += symbol * num
			num = 0
			hasLeading = false
		} else if str[i] == '+' {
			constants += symbol * num
			num = 0
			symbol = 1
			hasLeading = false
		} else if str[i] == '-' {
			constants += symbol * num
			symbol = -1
			num = 0
			hasLeading = false
		} else {
			hasLeading = true
			num = num*10 + int(str[i]-'0')
		}
	}

	if num > 0 {
		constants += symbol * num
	}

	return coefficient, constants
}

func main() {
	//fmt.Println(solveEquation("x+5-3+x=6+x-2"))
	//fmt.Println(solveEquation("x=x"))
	//fmt.Println(solveEquation("2x=x"))
	//fmt.Println(solveEquation("2x+3x-6x=x+2"))
	//fmt.Println(solveEquation("0x=0"))
	fmt.Println(solveEquation("1+x=2"))
}
