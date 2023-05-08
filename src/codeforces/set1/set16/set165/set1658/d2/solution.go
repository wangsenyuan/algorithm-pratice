package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		l, r := readTwoNums(reader)
		C := readNNums(reader, r-l+1)
		res := solve(l, r, C)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

func solve1(l, r int, A []int) int {

	var find func(l int, r int, s map[int]bool) int

	find = func(l int, r int, s map[int]bool) int {
		if l&1 == 0 && r&1 == 1 {
			t := make(map[int]bool)
			for k := range s {
				t[k>>1] = true
			}
			return find(l>>1, r>>1, t) << 1
		}
		for v := range s {
			if !s[v^1] {
				ans := v
				if l&1 == 0 {
					ans ^= r
				} else {
					ans ^= l
				}
				ok := true
				for x := range s {
					if x^ans < l || x^ans > r {
						ok = false
						break
					}
				}
				if ok {
					return ans
				}
			}
		}

		return 0
	}

	s := make(map[int]bool)

	for _, num := range A {
		s[num] = true
	}

	return find(l, r, s)
}

func solve(l, r int, A []int) int {
	root := new(Node)
	for _, num := range A {
		AddNum(root, num)
	}

	n := r - l + 1

	check := func(x int) bool {
		a, b := root, root
		var u, v int
		for i := H - 1; i >= 0; i-- {
			d := (x >> i) & 1
			if a.children[d^1] != nil {
				u |= 1 << i
				a = a.children[d^1]
			} else {
				a = a.children[d]
			}

			if b.children[d] != nil {
				b = b.children[d]
			} else {
				v |= 1 << i
				b = b.children[1^d]
			}
		}
		return u == r && v == l
	}

	for i := 0; i < n; i++ {
		x := A[i] ^ l
		if check(x) {
			return x
		}
	}
	return 0
}

type Node struct {
	children [2]*Node
}

const H = 17

func AddNum(root *Node, num int) {
	cur := root
	for i := H - 1; i >= 0; i-- {
		x := (num >> i) & 1
		if cur.children[x] == nil {
			cur.children[x] = new(Node)
		}
		cur = cur.children[x]
	}
}
