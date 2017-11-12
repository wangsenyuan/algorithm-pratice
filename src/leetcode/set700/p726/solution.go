package p726

import (
	"bytes"
	"sort"
	"strconv"
)

func countOfAtoms(formula string) string {
	if len(formula) == 0 {
		return ""
	}

	count := make(map[string]int)

	i := 0
	for i < len(formula) {
		if formula[i] == '(' {
			k := findPair(formula, i)
			subRes := countOfAtoms(formula[i+1 : k])
			j := k + 1
			for j < len(formula) && formula[j] >= '0' && formula[j] <= '9' {
				j++
			}
			num := 1
			if j > k+1 {
				tmp, _ := strconv.Atoi(formula[k+1 : j])
				num = tmp
			}

			a := 0
			for a < len(subRes) {
				c := a + 1
				for c < len(subRes) && subRes[c] >= 'a' && subRes[c] <= 'z' {
					c++
				}

				x := subRes[a:c]

				b := c
				for b < len(subRes) && subRes[b] >= '0' && subRes[b] <= '9' {
					b++
				}
				y := 1
				if b > c {
					tmp, _ := strconv.Atoi(subRes[c:b])
					y = tmp
				}
				y *= num
				count[x] += y
				a = b
			}

			i = j
		} else if formula[i] >= 'A' && formula[i] <= 'Z' {
			k := i + 1
			for k < len(formula) && formula[k] >= 'a' && formula[k] <= 'z' {
				k++
			}
			x := formula[i:k]
			j := k
			for j < len(formula) && formula[j] >= '0' && formula[j] <= '9' {
				j++
			}
			if j > k {
				num, _ := strconv.Atoi(formula[k:j])
				count[x] += num
			} else {
				count[x]++
			}
			i = j
		}
	}

	names := make([]string, 0, len(count))
	for name := range count {
		names = append(names, name)
	}

	sort.Strings(names)

	var buf bytes.Buffer
	for _, name := range names {
		if count[name] == 0 {
			continue
		}
		buf.WriteString(name)
		if count[name] > 1 {
			buf.WriteString(strconv.Itoa(count[name]))
		}
	}
	return buf.String()
}

func findPair(str string, start int) int {
	level := 0
	for i := start; i < len(str); i++ {
		if str[i] == '(' {
			level++
		} else if str[i] == ')' {
			level--
		}
		if level == 0 {
			return i
		}
	}
	return -1
}
