package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		s, _ := reader.ReadBytes('\n')
		var cnt int
		pos := readInt(s, 0, &cnt)
		a[i] = make([]int, cnt)
		for j := 0; j < cnt; j++ {
			pos = readInt(s, pos+1, &a[i][j])
		}
	}
	m := readNum(reader)
	bans := make([][]int, m)
	for i := 0; i < m; i++ {
		bans[i] = readNNums(reader, n)
	}
	res := solve(n, a, bans)

	s := fmt.Sprintf("%v", res)

	fmt.Println(s[1 : len(s)-1])
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

func solve(n int, a [][]int, bans [][]int) []int {
	base := make([]Hash, n)
	base[n-1] = NewHash(1)
	for i := n - 1; i > 0; i-- {
		base[i-1] = base[i].Mul(NewHash(len(a[i])))
	}

	// 1234
	//1000 100 10 1

	getKey := func(b []int) Hash {
		var res Hash
		for i := n - 1; i >= 0; i-- {
			res = res.Add(base[i].Mul(NewHash(b[i])))
		}

		return res
	}

	mem := make(map[Hash]bool)

	for _, b := range bans {
		for i := 0; i < n; i++ {
			b[i]--
		}
		mem[getKey(b)] = true
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = len(a[i]) - 1
	}

	output := func(arr []int) []int {
		for i := 0; i < n; i++ {
			arr[i]++
		}
		return arr
	}

	if !mem[getKey(arr)] {
		return output(arr)
	}

	var best int
	ans := make([]int, n)
	for _, b := range bans {
		var sum int
		for i := 0; i < n; i++ {
			sum += a[i][b[i]]
		}
		key := getKey(b)

		for i := 0; i < n; i++ {
			if b[i] > 0 {
				nk := key.Sub(base[i])
				nsum := sum - a[i][b[i]] + a[i][b[i]-1]
				if nsum > best && !mem[nk] {
					best = nsum
					copy(ans, b)
					ans[i]--
				}
			}
		}
	}

	return output(ans)
}

var MOD = [...]uint64{10000000007, 100000000009}

type Hash struct {
	h [2]uint64
}

func NewHash(x int) Hash {
	h := [2]uint64{uint64(x) % MOD[0], uint64(x) % MOD[1]}
	return Hash{h}
}

func (this Hash) Sub(that Hash) Hash {
	h := [2]uint64{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + MOD[i] - that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) Add(that Hash) Hash {
	h := [2]uint64{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) AddInt(x uint64) Hash {
	h := [2]uint64{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + x%MOD[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) Mul(that Hash) Hash {
	h := [2]uint64{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] * that.h[i]) % MOD[i]
	}
	return Hash{h}
}
