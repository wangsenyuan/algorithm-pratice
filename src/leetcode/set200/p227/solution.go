package main

import "fmt"

func main() {
	fmt.Println(calculate("161+6033/2/69-11*2-12*6+51*40"))
}

func calculate(s string) int {
	nums := make([]int, 10)
	p0 := 0
	ops := make([]byte, 10)
	p1 := 0

	i := 0

	for i < len(s) {
		b := s[i]
		if b == ' ' {
			i++
			continue
		}

		if b == '+' || b == '-' || b == '*' || b == '/' {
			for p1 > 0 && prefer(ops[p1-1], b) {
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

		num, j := readNum(s, i)
		nums = extendNum(nums, p0, num)
		p0++
		i = j
	}

	for p1 > 0 {
		num := doCalculate(nums, p0, ops, p1)
		p0 -= 2
		p1--
		nums = extendNum(nums, p0, num)
		p0++
	}

	return nums[0]
}

func prefer(a, b byte) bool {
	if a == '*' || a == '/' {
		return true
	}

	if b == '+' || b == '-' {
		return true
	}

	return false
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
