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

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		s := readString(reader)
		res := solve(k, s[:n])
		buf.WriteString(res)
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve(k int, s string) string {
	arr := []byte(s)
	sort.Sort(Bytes(arr))
	n := len(s)
	if k >= 26 {
		return string(arr)
	}
	S := []byte(s)
	var p int
	for i := 0; i < k-2; i++ {
		for j := 0; j < n && p < n; j++ {
			if S[j] == arr[p] {
				p++
				S[j] = '#'
			}
		}
	}

	if k > 1 {
		a := make([]bool, n)
		var pr byte = 'z'
		for i := n - 1; i >= 0; i-- {
			if S[i] != '#' {
				if S[i] <= pr {
					pr = S[i]
					a[i] = true
				}
			}
		}
		var y, z byte = '#', '#'
		for i := 0; i < n; i++ {
			if S[i] != '#' {
				if !a[i] {
					if y == '#' {
						y = S[i]
					} else if S[i] != y {
						z = S[i]
						break
					}
				}
			}
		}

		if z != '#' && z < y {
			y--
		}
		if y != '#' {
			for i := 0; i < n; i++ {
				if a[i] {
					if S[i] <= y {
						arr[p] = S[i]
						p++
						S[i] = '#'
					}
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if S[i] != '#' {
			arr[p] = S[i]
			p++
			S[i] = '#'
		}
	}
	return string(arr)
}

type Bytes []byte

func (this Bytes) Len() int {
	return len(this)
}

func (this Bytes) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Bytes) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
