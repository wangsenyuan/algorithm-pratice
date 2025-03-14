package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func process(reader *bufio.Reader) string {
	s := readString(reader)
	k := readNum(reader)
	return solve(s, k)
}

func solve(s string, k int) string {
	n := len(s)

	good := make([][]bool, n)
	for i := range n {
		good[i] = make([]bool, n)
	}

	for r := 0; r < n; r++ {
		for l := r; l >= 0; l-- {
			if l == r {
				good[l][r] = true
			} else {
				good[l][r] = s[l] == s[r] && (l+2 > r-2 || good[l+2][r-2])
			}
		}
	}

	tr := new(Trie)

	for l := 0; l < n; l++ {
		var sum int
		for r := l; r < n; r++ {
			if good[l][r] {
				sum++
			}
		}
		tmp := tr
		for r := l; r < n; r++ {
			tmp = tmp.Add(s[r])
			tmp.sum += sum
			if good[l][r] {
				tmp.cnt++
				sum--
			}
		}
	}

	var buf []byte

	tmp := tr

	for tmp != nil && k > 0 {
		if tmp.children[0] != nil && tmp.children[0].sum >= k {
			// go a
			buf = append(buf, 'a')
			tmp = tmp.children[0]
		} else {
			// go b
			buf = append(buf, 'b')
			if tmp.children[0] != nil {
				k -= tmp.children[0].sum
			}
			tmp = tmp.children[1]
		}
		k -= tmp.cnt
	}

	return string(buf)
}

type Trie struct {
	children [2]*Trie
	sum      int
	cnt      int
}

func (tr *Trie) Add(c byte) *Trie {
	x := int(c - 'a')
	if tr.children[x] == nil {
		tr.children[x] = new(Trie)
	}
	return tr.children[x]
}
