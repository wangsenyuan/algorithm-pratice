package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))

	for _, lyric := range res {
		buf.WriteString(lyric[0])
		buf.WriteByte('\n')
		buf.WriteString(lyric[1])
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) [][]string {
	// 使用reader读取n条string
	n := readNum(reader)
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = readString(reader)
	}
	return solve(words)
}

func solve(words []string) [][]string {

	states := make(map[state][]string)

	for _, w := range words {
		s := convert(w)
		states[s] = append(states[s], w)
	}

	// 能否组成x个满足条件的对仗句子
	check := func(x int) bool {
		cnt := make(map[int]int)
		var y int
		for s, vs := range states {
			v := len(vs)
			u := min(x-y, v/2)
			for w := 0; 2*u+w < v; w++ {
				cnt[s.cnt]++
			}
			y += u
		}
		if y < x {
			return false
		}
		// y == x
		for _, v := range cnt {
			x -= v / 2
		}
		return x <= 0
	}
	l, r := 0, len(words)/2
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	// 可以有这么多个对仗句子
	cnt := make(map[int][]string)
	x := r - 1

	if x <= 0 {
		return nil
	}

	var second [][]string

	var y, z int
	for s, vs := range states {
		u := min(x-y, len(vs)/2)
		for i := 0; i < 2*u; i += 2 {
			tmp := []string{vs[i], vs[i+1]}
			second = append(second, tmp)
		}
		for i := 2 * u; i < len(vs); i++ {
			cnt[s.cnt] = append(cnt[s.cnt], vs[i])
		}
		y += u
	}

	var first [][]string

	for _, vs := range cnt {
		u := min(x-z, len(vs)/2)
		for i := 0; i < 2*u; i += 2 {
			tmp := []string{vs[i], vs[i+1]}
			first = append(first, tmp)
		}
		z += u
	}

	ans := make([][]string, 0, r)

	for i := 0; i < x; i++ {
		a := first[i]
		b := second[i]
		tmp := []string{a[0] + " " + b[0], a[1] + " " + b[1]}
		ans = append(ans, tmp)
	}

	return ans
}

type state struct {
	cnt  int
	last byte
}

const vowels = "aeiou"

func convert(w string) state {
	var s state
	for i := 0; i < len(w); i++ {
		if strings.Contains(vowels, string(w[i])) {
			s.cnt++
			s.last = w[i]
		}
	}
	return s
}
