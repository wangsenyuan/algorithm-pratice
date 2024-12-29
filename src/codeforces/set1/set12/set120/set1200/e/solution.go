package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
	s := readString(reader)
	words := strings.Split(s, " ")
	res := solve(words)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func solve(words []string) string {
	var m int
	for _, w := range words {
		m += len(w)
	}
	next := make([]int, m)

	res := []byte(words[0])

	kmp := func(a string, b []byte) int {
		n := len(a) + len(b)
		clear(next[:n])
		get := func(p int) byte {
			if p < len(a) {
				return a[p]
			}
			return b[p-len(a)]
		}

		for i := 1; i < len(a)+len(b); i++ {
			j := next[i-1]
			for j > 0 && get(i) != get(j) {
				j = next[j-1]
			}
			if get(i) == get(j) {
				j++
			}
			next[i] = j
		}

		return n
	}

	for i := 1; i < len(words); i++ {
		w := words[i]
		x := len(res)
		y := len(w)
		var b []byte
		if x > y {
			b = res[x-y-1:]
		} else {
			b = res
		}
		n := kmp(w, b)
		// l := min(next[n-1], y, x)
		l := next[n-1]
		for l > x || l > y {
			l = next[l-1]
		}

		res = append(res, w[l:]...)
	}

	return string(res)
}

func solve1(words []string) string {
	var m int
	for _, w := range words {
		m += len(w)
	}
	bases := make([]Key, m+1)
	bases[0] = NewKey(1)

	for i := 1; i <= m; i++ {
		bases[i] = bases[i-1].Mul(62)
	}

	var buf []byte
	var arr []Key
	arr = append(arr, NewKey(0))

	add := func(w string) {
		key := arr[len(arr)-1]
		for j := 0; j < len(w); j++ {
			key = key.Mul(62).Add(convert(w[j]))
			arr = append(arr, key)
			buf = append(buf, w[j])
		}
	}

	add(words[0])

	for i := 1; i < len(words); i++ {
		w := words[i]
		var key Key
		n := len(arr) - 1
		pos := n - 1
		at := n - 1
		for j := 0; j < len(w) && pos >= 0; j++ {
			key = key.Mul(62).Add(convert(w[j]))
			tmp2 := arr[n].Sub(arr[pos].Mul2(bases[n-pos]))
			if key == tmp2 {
				at = pos - 1
			}
			pos--
		}
		// whole new words
		arr = arr[:at+2]
		buf = buf[:at+1]
		add(w)
	}

	return string(buf)
}

func convert(x byte) int {
	if x >= '0' && x <= '9' {
		return int(x - '0')
	}
	if x >= 'A' && x <= 'Z' {
		return 10 + int(x-'A')
	}
	return 36 + int(x-'a')
}

const MOD1 = 1e9 + 7
const MOD2 = 1e9 + 9
const MOD3 = 1e9 + 23

var mods = [3]int{MOD1, MOD2, MOD3}

type Key struct {
	vals [3]int
}

func NewKey(v int) Key {
	var vals [3]int
	for i := 0; i < len(vals); i++ {
		vals[i] = v % mods[i]
	}
	return Key{vals}
}

func (this Key) Mul(x int) Key {
	var vals [3]int
	for i := 0; i < len(vals); i++ {
		vals[i] = this.vals[i] * x % mods[i]
	}

	return Key{vals}
}

func (this Key) Add(x int) Key {
	var vals [3]int
	for i := 0; i < len(vals); i++ {
		vals[i] = (this.vals[i] + x) % mods[i]
	}

	return Key{vals}
}

func (this Key) Sub(that Key) Key {
	var vals [3]int
	for i := 0; i < len(vals); i++ {
		vals[i] = (this.vals[i] - that.vals[i] + mods[i]) % mods[i]
	}
	return Key{vals}
}

func (this Key) Mul2(that Key) Key {
	var vals [3]int
	for i := 0; i < len(vals); i++ {
		vals[i] = this.vals[i] * that.vals[i] % mods[i]
	}
	return Key{vals}
}
