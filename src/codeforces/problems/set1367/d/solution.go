package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		s, _ := reader.ReadString('\n')
		s = s[:len(s)-1]
		m := readNum(reader)
		B := readNNums(reader, m)

		fmt.Println(solve(s, m, B))
	}
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

func solve(s string, m int, B []int) string {
	mem := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		mem[s[i]]++
	}

	xx := make([]byte, 0, len(mem))

	for x := range mem {
		xx = append(xx, x)
	}

	sort.Sort(Bytes(xx))

	res := make([]byte, m)

	for i := 0; i < m; i++ {
		res[i] = ' '
	}
	zs := make([]int, m)

	for i := 0; i < len(xx); i++ {
		x := xx[i]

		// find 0
		var p int
		var sum int
		for j := 0; j < m; j++ {
			if B[j] == 0 && res[j] == ' ' {
				zs[p] = j
				p++
				sum += j
			}
		}

		if mem[x] < p {
			continue
		}

		for j := 0; j < p; j++ {
			res[zs[j]] = x
		}

		var sum2 int
		var k int
		for j := 0; j < m; j++ {
			if B[j] == 0 {
				continue
			}
			for k < p && zs[k] < j {
				sum2 += zs[k]
				k++
			}
			B[j] -= (k*j - sum2)
			B[j] -= (sum - sum2 - (p-k)*j)
		}
	}

	return string(res)
}

type Bytes []byte

func (bs Bytes) Len() int {
	return len(bs)
}

func (bs Bytes) Less(i, j int) bool {
	return bs[i] > bs[j]
}

func (bs Bytes) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}
