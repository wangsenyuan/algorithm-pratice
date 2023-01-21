package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)

	lamps := make([][]int, n)

	for i := 0; i < n; i++ {
		lamps[i] = readNNums(reader, 2)
	}

	fmt.Println(solve(n, k, lamps))
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

const MOD = 998244353

const MAX_N = 300010

var fact []int64

var ifact []int64

func init() {
	fact = make([]int64, MAX_N+1)
	ifact = make([]int64, MAX_N+1)

	fact[0] = 1
	for i := 1; i <= MAX_N; i++ {
		fact[i] = (fact[i-1] * int64(i)) % MOD
	}

	ifact[MAX_N] = inverse(fact[MAX_N])

	for i := MAX_N; i > 0; i-- {
		ifact[i-1] = (int64(i) * ifact[i]) % MOD
	}
}

func nCr(n int, r int) int64 {
	if n < r {
		return 0
	}
	res := (ifact[r] * ifact[n-r]) % MOD
	res *= fact[n]
	return res % MOD
}

func inverse(q int64) int64 {
	return pow(q, MOD-2)
}

func pow(a, b int64) int64 {
	r := int64(1)

	for b > 0 {
		if b&1 == 1 {
			r = (r * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return r
}

func solve(n, k int, lamps [][]int) int {
	events := make([]int, 0, 2*n)

	for i := 0; i < n; i++ {
		cur := lamps[i]
		events = append(events, cur[0]*2)
		events = append(events, cur[1]*2+1)
	}
	var res int64
	sort.Ints(events)
	var on int
	for i := 0; i < len(events); {
		j := i
		for i < len(events) && events[j] == events[i] {
			i++
		}
		added := i - j
		if events[j]%2 == 0 {
			// on events
			res += nCr(on+added, k)
			res %= MOD
			res += MOD - nCr(on, k)
			res %= MOD
			on += added
		} else {
			on -= added
		}
	}

	return int(res)
}

func solve1(n, k int, lamps [][]int) int {
	pp := make([]int, 0, n*2)

	for _, lamp := range lamps {
		l, r := lamp[0], lamp[1]
		pp = append(pp, l)
		pp = append(pp, r)
	}
	sort.Ints(pp)
	ii := make(map[int]int)
	var p int
	for i := 0; i < len(pp); i++ {
		if _, found := ii[pp[i]]; !found {
			ii[pp[i]] = p
			p++
		}
	}
	ls := make([]Lamp, n)
	for i := 0; i < n; i++ {
		ls[i] = Lamp{ii[lamps[i][0]], ii[lamps[i][1]]}
	}

	sort.Sort(Lamps(ls))

	sum := make([]int, p+1)
	var res int64
	for i, j := 0, 0; i <= p && j < n; i++ {
		if i > 0 {
			sum[i] += sum[i-1]
		}
		if i < ls[j].on {
			continue
		}
		// i == ls[j].on
		for j < n && i == ls[j].on {
			sum[i]++
			sum[ls[j].off+1]--
			if sum[i] >= k {
				// take another k - 1 besides j
				res += nCr(sum[i]-1, k-1)
				res %= MOD
			}
			j++
		}
	}

	return int(res)
}

type Lamp struct {
	on  int
	off int
}

type Lamps []Lamp

func (ls Lamps) Len() int {
	return len(ls)
}

func (ls Lamps) Less(i, j int) bool {
	return ls[i].on < ls[j].on || ls[i].on == ls[j].on && ls[i].off > ls[j].off
}

func (ls Lamps) Swap(i, j int) {
	ls[i], ls[j] = ls[j], ls[i]
}
