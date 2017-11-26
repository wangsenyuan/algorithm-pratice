package p736

type Scope struct {
	parent  *Scope
	current map[string]int
}

func newScope(parent *Scope) Scope {
	return Scope{parent, make(map[string]int)}
}

func (scope Scope) find(vn string) (int, bool) {
	if vv, found := scope.current[vn]; found {
		return vv, true
	}

	if scope.parent == nil {
		return -1, false
	}
	return scope.parent.find(vn)
}

func (scope *Scope) add(vn string, vv int) {
	scope.current[vn] = vv
}

func process(expr string, parent *Scope) int {
	n := len(expr)
	// get rid of the out most ()
	expr = expr[1 : n-1]
	op := expr[:3]
	switch op {
	case "let":
		return let(expr[4:], newScope(parent))
	case "add":
		return add(expr[4:], newScope(parent))
	default:
		return mul(expr[5:], newScope(parent))
	}
}

func eval(expr string, scope Scope) int {
	if expr[0] == '(' {
		return process(expr, &scope)
	}

	if expr[0] >= 'a' && expr[0] <= 'z' {
		res, _ := scope.find(expr)
		return res
	}

	var tmp int
	sign := 1
	if expr[0] == '-' {
		sign = -1
		expr = expr[1:]
	}
	for i := 0; i < len(expr); i++ {
		tmp = tmp*10 + int(expr[i]-'0')
	}
	return tmp * sign
}

func pair(expr string, start int) int {
	level := 0
	for i := start; i < len(expr); i++ {
		if expr[i] == '(' {
			level++
		} else if expr[i] == ')' {
			level--
		}
		if level == 0 {
			return i
		}
	}
	return -1
}

func let(expr string, scope Scope) int {
	i, k := 0, 0
	var vn string
	for i < len(expr) {
		if k%2 == 0 {
			// read the v part
			if expr[i] == '(' {
				return eval(expr[i:], scope)
			}
			j := i
			for i < len(expr) && expr[i] != ' ' {
				i++
			}
			if i == len(expr) {
				return eval(expr[j:], scope)
			}
			vn = expr[j:i]
			k++
		} else {
			// read the e part
			var res int
			if expr[i] == '(' {
				j := pair(expr, i)
				res = eval(expr[i:j+1], scope)
				// skip the right )
				i = j + 1
			} else {
				j := i
				for i < len(expr) && expr[i] != ' ' {
					i++
				}
				res = eval(expr[j:i], scope)
			}
			scope.add(vn, res)
			k++
		}
		//skip the blank
		i++
	}
	return -1
}

func calc(expr string, scope Scope, fn func(int, int) int) int {
	var left, right int
	if expr[0] == '(' {
		i := pair(expr, 0)
		left = eval(expr[:i+1], scope)
		// skip the ) and blank
		right = eval(expr[i+2:], scope)
	} else {
		i := 0
		for expr[i] != ' ' {
			i++
		}
		left = eval(expr[:i], scope)
		right = eval(expr[i+1:], scope)
	}
	return fn(left, right)
}
func add(expr string, scope Scope) int {
	return calc(expr, scope, func(left int, right int) int { return left + right })
}

func mul(expr string, scope Scope) int {
	return calc(expr, scope, func(left int, right int) int { return left * right })
}

func evaluate(expression string) int {
	return process(expression, nil)
}
