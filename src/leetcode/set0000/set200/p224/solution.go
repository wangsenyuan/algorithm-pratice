package main

import "fmt"

func main() {
	fmt.Println(calculate("1 + 2"))
	fmt.Println(calculate(" 2-1 + 2 "))
	fmt.Println(calculate(" (1+(4+5+2)-3)+(6+8) "))
	fmt.Println(calculate(" 0 "))
	fmt.Println(calculate(" 2 - (5 ) "))

}

func calculate(s string) int {
	s = normalize(s)
	nums := make([]int, 10)
	p0 := 0
	ops := make([]byte, 10)
	p1 := 0

	i := 0

	for i < len(s) {
		b := s[i]
		if b == '+' || b == '-' {
			for p1 > 0 {
				num := doCalculate(nums, p0, ops, p1)
				p0 -= 2
				p1--
				nums = extendNum(nums, p0, num)
				p0++
			}
			ops = extendOp(ops, p1, b)
			p1++
			i++
			continue
		}

		if b == '(' {
			j := pair(s, i, 0)
			num := calculate(s[i+1 : j])
			nums = extendNum(nums, p0, num)
			p0++
			i = j + 1
			continue
		}

		num, j := readNum(s, i)
		nums = extendNum(nums, p0, num)
		p0++
		i = j
	}

	if p1 > 0 {
		return doCalculate(nums, p0, ops, p1)
	}

	return nums[0]
}

func normalize(s string) string {
	buf := []byte(s)
	var n int
	for i := 0; i < len(s); i++ {
		if buf[i] == ' ' {
			continue
		}
		buf[n] = buf[i]
		n++
	}
	s = string(buf[:n])
	if s[0] == '-' {
		return "0" + s
	}
	return s
}

func readNum(s string, start int) (int, int) {
	num := 0
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return num, i
		}
		num = num*10 + int(s[i]-'0')
	}
	return num, len(s)
}

func pair(s string, i int, level int) int {
	if s[i] == ')' {
		if level == 1 {
			return i
		}
		return pair(s, i+1, level-1)
	}

	if s[i] == '(' {
		return pair(s, i+1, level+1)
	}

	return pair(s, i+1, level)
}

func extendNum(nums []int, sz int, num int) []int {
	if sz < len(nums) {
		nums[sz] = num
		return nums
	}
	return append(nums, num)
}

func extendOp(ops []byte, sz int, b byte) []byte {
	if sz < len(ops) {
		ops[sz] = b
		return ops
	}

	return append(ops, b)
}

func doCalculate(nums []int, p0 int, ops []byte, p1 int) int {
	a, b := nums[p0-2], nums[p0-1]
	op := ops[p1-1]

	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}
	panic("wroooooong")
}
