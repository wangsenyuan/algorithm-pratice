package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		nums := readNNums(reader, 4)
		s := readString(reader)
		res := solve(s, nums)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(s string, num []int) bool {
	n := len(s)
	a, b, c, d := num[0], num[1], num[2], num[3]

	if a+b+(c+d)*2 != n {
		return false
	}
	cnt := make([]int, 2)
	for i := 0; i < n; i++ {
		cnt[int(s[i]-'A')]++
	}

	// A的数量 = a + c + d
	if a+c+d != cnt[0] || b+c+d != cnt[1] {
		return false
	}

	// BAABBABBAA
	// BA AB B AB BA A
	// BA AB BA B BA A

	var ab []int
	var ba []int
	var other []int

	for i, j := 0, 0; i < n; i++ {
		if i+1 == n || s[i] == s[i+1] {
			l := i - j + 1
			if l%2 == 0 {
				if s[j] == 'A' {
					ab = append(ab, l)
				} else {
					ba = append(ba, l)
				}
			} else {
				other = append(other, l)
			}
			j = i + 1
		}
	}

	sort.Ints(ab)
	sort.Ints(ba)

	for i := 0; i < len(ab); i++ {
		x := ab[i]
		x /= 2
		if c < x {
			// c < x
			ab[i] = 2 * (x - c)
			c = 0
			break
		}
		ab[i] = 0
		c -= x
	}

	for i := 0; i < len(ba); i++ {
		x := ba[i]
		x /= 2
		if d < x {
			ba[i] = 2 * (x - d)
			d = 0
			break
		}
		ba[i] = 0
		d -= x
	}

	other = append(other, ab...)
	other = append(other, ba...)

	sort.Ints(other)

	for i := 0; i < len(other); i++ {
		if other[i] <= 1 {
			continue
		}
		// other[i] > 2
		k := (other[i] - 1) / 2
		// x + y <= k
		if c > 0 {
			x := min(c, k)
			c -= x
			k -= x
		}
		if d > 0 {
			x := min(d, k)
			d -= x
			k -= x
		}
	}

	return c == 0 && d == 0
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
