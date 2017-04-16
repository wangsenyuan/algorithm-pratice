package main

import "fmt"

func main() {

	fmt.Println(splitLoopedString([]string{"abc", "xyz"}))

	fmt.Println(splitLoopedString([]string{"lc", "evol", "cdy"})) //ylclovecd
}

func splitLoopedString(strs []string) string {
	mx := 'a'
	var mxi = make([]int, 0, len(strs))
	for i, str := range strs {
		for _, c := range str {
			if c > mx {
				mx = c
				mxi = make([]int, 0, len(strs))
				mxi = append(mxi, i)
			} else if c == mx {
				mxi = append(mxi, i)
			}
		}
	}

	n := len(strs)

	var process func(res string, i, end int) string
	process = func(res string, i, end int) string {
		if i == end {
			return res
		}
		str := strs[i]
		tmp := reverse(str)
		if tmp > str {
			return process(res+tmp, (i+1)%n, end)
		}
		return process(res+str, (i+1)%n, end)
	}

	var best = ""
	for _, i := range mxi {
		tmp := process("", (i+1)%n, i)
		a := splitAndConcat(tmp, strs[i], mx)
		b := splitAndConcat(tmp, reverse(strs[i]), mx)
		best = max(best, max(a, b))
	}
	return best
}
func splitAndConcat(a string, b string, x rune) string {
	var best = ""
	for i, y := range b {
		if y == x {
			tmp := b[i:] + a + b[:i]
			if best == "" || tmp > best {
				best = tmp
			}
		}
	}
	return best
}

func reverse(str string) string {
	tmp := []byte(str)
	i, j := 0, len(str)-1
	for i < j {
		tmp[i], tmp[j] = tmp[j], tmp[i]
		i++
		j--
	}
	return string(tmp)
}

func max(a, b string) string {
	if a == "" || a < b {
		return b
	}
	return a
}
