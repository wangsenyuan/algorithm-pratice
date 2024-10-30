package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	res := solve(k, a)
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

func solve(k int, a []int) int {
	var ans int

	tr := new(Trie)

	calc := func(num int) {
		node := tr
		var val int
		for d := H - 1; d >= 0 && node != nil; d-- {
			x := (num >> d) & 1
			if val|(1<<d) >= k {
				// 如果某个prefix的d位和当前位不一致，xor的结果超过了k
				// 那么就加入cnt，
				ans += node.children[x^1].Get()
				node = node.children[x]
			} else {
				node = node.children[x^1]
				val |= 1 << d
			}
		}
		if val >= k {
			ans += node.Get()
		}
	}
	var sum int
	tr.Add(sum, H-1)
	for _, num := range a {
		sum ^= num
		calc(sum)
		tr.Add(sum, H-1)
	}

	return ans
}

const H = 30

type Trie struct {
	children [2]*Trie
	cnt      int
}

func (tr *Trie) Get() int {
	if tr == nil {
		return 0
	}
	return tr.cnt
}

func (tr *Trie) Add(num int, d int) {
	tr.cnt++
	if d < 0 {
		return
	}
	x := (num >> d) & 1
	if tr.children[x] == nil {
		tr.children[x] = new(Trie)
	}
	tr.children[x].Add(num, d-1)
}
