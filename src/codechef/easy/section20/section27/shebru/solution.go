package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	ss := make([]string, n)
	for i := 0; i < n; i++ {
		ss[i] = readString(reader)
	}
	fmt.Println(solve(ss))
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

const MOD = 1000000007
const MAX = 100010

func update(bit []int, p int, v int) {
	for p < len(bit) {
		add(&bit[p], v)
		p += p & -p
	}
}

func get(bit []int, p int) int {
	var res int
	for p > 0 {
		add(&res, bit[p])
		p -= p & -p
	}
	return res
}

func solve(ss []string) int {
	n := len(ss)

	v := make([][]int, 26)

	for i := 0; i < n; i++ {
		s := ss[i]
		x := int(s[0] - 'a')
		if len(v[x]) == 0 {
			v[x] = make([]int, 0, 1)
		}
		v[x] = append(v[x], len(s))
	}

	var ans int
	x := make([]int, MAX+1)
	y := make([]int, MAX+1)
	for i := 0; i < 26; i++ {
		for j := 0; j < MAX; j++ {
			x[j] = 0
			y[j] = 0
		}
		for _, l := range v[i] {
			tmp := get(x, l-1)
			add(&ans, tmp)
			update(y, l, (tmp+1)%MOD)
			tmp = get(y, MAX)
			sub(&tmp, get(y, l))
			add(&ans, tmp)
			update(x, l, (tmp+1)%MOD)
		}
	}

	return ans
}

func add(a *int, b int) {
	*a += b
	if *a >= MOD {
		*a -= MOD
	}
}

func sub(a *int, b int) {
	add(a, MOD-b)
}
