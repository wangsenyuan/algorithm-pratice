package p1106

func parseBoolExpr(expression string) bool {
	n := len(expression)
	ps := make([]int, n)

	stack := make([]int, n)
	var p int

	for i := 0; i < n; i++ {
		if expression[i] == '(' {
			stack[p] = i
			p++
		} else if expression[i] == ')' {
			ps[stack[p-1]] = i
			p--
		} else {
			ps[i] = -1
		}
	}

	var loop func(start int, end int) bool
	var andExpr func(start, end int) bool
	var orExpr func(start, end int) bool
	loop = func(start int, end int) bool {
		if ps[start] == end {
			return loop(start+1, end-1)
		}

		if expression[start] == 't' {
			return true
		}

		if expression[start] == 'f' {
			return false
		}

		if expression[start] == '!' {
			return !loop(start+1, end)
		}

		if expression[start] == '&' {
			return andExpr(start, end)
		}
		return orExpr(start, end)
	}

	andExpr = func(start, end int) bool {
		i := start + 2
		j := start + 2
		for i <= end {
			if i == end || expression[i] == ',' {
				tmp := loop(j, i-1)
				if !tmp {
					return false
				}
				j = i + 1
			} else if expression[i] == '(' {
				i = ps[i]
			}
			i++
		}
		return true
	}

	orExpr = func(start, end int) bool {
		i := start + 2
		j := start + 2
		for i <= end {
			if i == end || expression[i] == ',' {
				tmp := loop(j, i-1)
				if tmp {
					return true
				}
				j = i + 1
			} else if expression[i] == '(' {
				i = ps[i]
			}
			i++
		}
		return false
	}

	return loop(0, n-1)
}
