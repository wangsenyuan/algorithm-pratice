package p772

import "strconv"

func calculate(s string) int {
	s = eliminateSpaces(s)

	return int(eval(s))
}

func eval(s string) int64 {
	toks := tokenize(s)
	nums := make([]int64, len(toks))
	ops := make([]string, len(toks))
	var x, y int

	for i := 0; i < len(toks); i++ {
		if toks[i] == "+" || toks[i] == "-" || toks[i] == "*" || toks[i] == "/" {
			for y > 0 && prefer(ops[y-1], toks[i]) {
				tmp := calc(nums[x-2], nums[x-1], ops[y-1])
				x -= 2
				nums[x] = tmp
				x++
				y--
			}
			ops[y] = toks[i]
			y++
		} else if toks[i][0] == '(' {
			nums[x] = eval(toks[i][1 : len(toks[i])-1])
			x++
		} else {
			v, _ := strconv.Atoi(toks[i])
			nums[x] = int64(v)
			x++
		}
	}
	for y > 0 {
		tmp := calc(nums[x-2], nums[x-1], ops[y-1])
		x -= 2
		nums[x] = tmp
		x++
		y--
	}
	return nums[0]
}

func calc(a, b int64, op string) int64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		return a / b
	}
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

func eliminateSpaces(s string) string {
	b := []byte(s)
	var sz int
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			continue
		}
		b[sz] = b[i]
		sz++
	}
	return string(b[:sz])
}

func tokenize(s string) []string {
	res := make([]string, 0, 10)
	for i, j := 0, 0; i <= len(s); {
		if i == len(s) {
			if i > j {
				res = append(res, s[j:i])
			}
			i++
		} else if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			if i > j {
				res = append(res, s[j:i])
			}
			res = append(res, s[i:i+1])
			i++
			j = i
		} else if s[i] == '(' {
			lvl := 0
			k := i
			for k < len(s) {
				if s[k] == '(' {
					lvl++
				} else if s[k] == ')' {
					lvl--
				}
				k++
				if lvl == 0 {
					break
				}
			}
			res = append(res, s[i:k])
			i = k
			j = i
		} else {
			i++
		}
	}
	return res
}
