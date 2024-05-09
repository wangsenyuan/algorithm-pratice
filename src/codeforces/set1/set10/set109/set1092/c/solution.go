package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	words := make([]string, 2*n-2)
	for i := 0; i < 2*n-2; i++ {
		words[i] = readString(reader)
	}

	res := solve(n, words)

	fmt.Print(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func solve(n int, words []string) string {
	id := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		return len(words[id[i]]) < len(words[id[j]])
	})

	check := func(s string) bool {
		// 同一个长度的，只能一个是prefix，一个是suffix
		for i := 0; i < len(words); i += 2 {
			a := words[id[i]]
			b := words[id[i+1]]
			if isPrefix(s, a) && isSuffix(s, b) || isPrefix(s, b) && isSuffix(s, a) {
				continue
			}
			return false
		}
		return true
	}

	mark := func(s string) string {
		res := make([]byte, len(words))
		for i := 0; i < len(words); i += 2 {
			a := words[id[i]]
			if isPrefix(s, a) {
				res[id[i]] = 'P'
				res[id[i+1]] = 'S'
			} else {
				res[id[i]] = 'S'
				res[id[i+1]] = 'P'
			}
		}
		return string(res)
	}

	first := id[:2]
	last := id[len(id)-2:]

	for _, a := range first {
		for _, b := range last {
			if check(words[a] + words[b]) {
				return mark(words[a] + words[b])
			}
		}
	}

	return ""
}

func isSuffix(s string, a string) bool {
	for i := 0; i < len(a); i++ {
		if s[len(s)-1-i] != a[len(a)-1-i] {
			return false
		}
	}
	return true
}

func isPrefix(s string, a string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != s[i] {
			return false
		}
	}
	return true
}
