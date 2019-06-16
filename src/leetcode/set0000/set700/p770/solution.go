package p770

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func basicCalculatorIV(expression string, evalvars []string, evalints []int) []string {
	ctx := make(map[string]int)
	for i := 0; i < len(evalvars); i++ {
		ctx[evalvars[i]] = evalints[i]
	}

	expression = elminateSpace(expression)
	poly := eval(expression)
	poly = poly.Evaulate(ctx)

	return toList(poly)
}

func toList(poly Poly) []string {
	terms := poly.terms
	sort.Sort(SS(terms))

	res := make([]string, 0, len(poly.terms))

	for i := 0; i < len(terms); i++ {
		term := terms[i]
		s := asStringKey(term)
		c := poly.count[s]
		if c == 0 {
			continue
		}

		if len(s) == 0 {
			res = append(res, strconv.Itoa(c))
		} else {
			res = append(res, fmt.Sprintf("%d*%s", c, s))
		}
	}
	return res
}

type SS [][]string

func (ss SS) Len() int {
	return len(ss)
}

func compareSlice(a, b []string) bool {
	for k := 0; k < len(a); k++ {
		if a[k] < b[k] {
			return true
		} else if a[k] > b[k] {
			return false
		}
	}
	return false
}

func (ss SS) Less(i, j int) bool {
	if len(ss[i]) > len(ss[j]) {
		return true
	}
	if len(ss[i]) == len(ss[j]) {
		return compareSlice(ss[i], ss[j])
	}
	return false
}

func (ss SS) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func eval(expr string) Poly {
	toks := tokenize(expr)
	stack := make([]Poly, len(toks))
	x := 0
	ops := make([]string, len(toks))
	y := 0
	for i := 0; i < len(toks); i++ {
		tok := toks[i]
		if tok == "+" || tok == "-" || tok == "*" {
			for y > 0 && prefer(ops[y-1], tok) {
				tmp := combine(stack[x-2], ops[y-1], stack[x-1])
				x -= 2
				stack[x] = tmp
				x++
				y--
			}
			ops[y] = tok
			y++
		} else {
			stack[x] = parse(tok)
			x++
		}
	}

	for y > 0 {
		tmp := combine(stack[x-2], ops[y-1], stack[x-1])
		x -= 2
		stack[x] = tmp
		x++
		y--
	}
	return stack[0]
}

func parse(s string) Poly {
	if s[0] == '(' {
		return eval(s[1 : len(s)-1])
	}
	res := Poly{make(map[string]int), make([][]string, 0, 10)}

	if s[0] >= 'a' && s[0] <= 'z' {
		term := []string{s}
		res.update(term, s, 1)
	} else {
		term := []string{}
		v, _ := strconv.Atoi(s)
		res.update(term, "", v)
	}
	return res
}

func combine(a Poly, op string, b Poly) Poly {
	if op == "+" {
		return a.Add(b)
	}
	if op == "-" {
		return a.Sub(b)
	}
	return a.Mul(b)
}

func prefer(x, y string) bool {
	if x == "*" {
		return true
	}
	if y == "+" || y == "-" {
		return true
	}
	return false
}

func tokenize(s string) []string {
	res := make([]string, 0, 10)
	for i, j := 0, 0; i <= len(s); {
		if i == len(s) {
			if i > j {
				res = append(res, s[j:i])
			}
			break
		}
		if s[i] == '+' || s[i] == '-' || s[i] == '*' {
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

func elminateSpace(s string) string {
	b := []byte(s)
	sz := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		b[sz] = b[i]
		sz++
	}
	return string(b[:sz])
}

type Poly struct {
	count map[string]int
	terms [][]string
}

func asStringKey(term []string) string {
	if len(term) == 0 {
		return ""
	}
	return strings.Join(term, "*")
}

func (this *Poly) update(term []string, key string, val int) {
	if old, found := this.count[key]; found {
		this.count[key] = old + val
	} else {
		this.count[key] = val
		this.terms = append(this.terms, term)
	}
}

func (this Poly) Add(that Poly) Poly {
	res := Poly{make(map[string]int), make([][]string, 0, 10)}

	for _, term := range this.terms {
		key := asStringKey(term)
		res.update(term, key, this.count[key])
	}

	for _, term := range that.terms {
		key := asStringKey(term)
		res.update(term, key, that.count[key])
	}
	return res
}

func (this Poly) Sub(that Poly) Poly {
	res := Poly{make(map[string]int), make([][]string, 0, 10)}

	for _, term := range this.terms {
		key := asStringKey(term)
		res.update(term, key, this.count[key])
	}

	for _, term := range that.terms {
		key := asStringKey(term)
		res.update(term, key, -that.count[key])
	}
	return res
}

func (this Poly) Mul(that Poly) Poly {
	res := Poly{make(map[string]int), make([][]string, 0, 10)}
	for _, k1 := range this.terms {
		for _, k2 := range that.terms {
			k := make([]string, len(k1)+len(k2))
			copy(k, k1)
			copy(k[len(k1):], k2)
			sort.Strings(k)
			res.update(k, asStringKey(k), this.count[asStringKey(k1)]*that.count[asStringKey(k2)])
		}
	}
	return res
}

func (this Poly) Evaulate(ctx map[string]int) Poly {
	res := Poly{make(map[string]int), make([][]string, 0, 10)}

	for _, term := range this.terms {
		c := this.count[asStringKey(term)]
		free := make([]string, 0, len(term))
		for _, tok := range term {
			if v, found := ctx[tok]; found {
				c *= v
			} else {
				free = append(free, tok)
			}
		}
		res.update(free, asStringKey(free), c)
	}

	return res
}
