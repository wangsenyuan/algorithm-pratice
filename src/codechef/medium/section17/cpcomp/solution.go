package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	res := solve(A)
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(A []int) int {
	mx := max_of(A)
	lf := make([]int, mx+1)
	var primes []int

	for i := 2; i <= mx; i++ {
		if lf[i] == 0 {
			lf[i] = i
			primes = append(primes, i)
			if mx/i < i {
				continue
			}
			for j := i * i; j <= mx; j += i {
				if lf[j] == 0 {
					lf[j] = i
				}
			}
		}
	}

	simplify := func(num int) int {
		res := 1
		for num > 1 {
			x := lf[num]
			res *= x
			for num%x == 0 {
				num /= x
			}
		}
		return res
	}

	n := len(A)
	freq := make(map[int]int)
	for i := 0; i < n; i++ {
		num := simplify(A[i])
		freq[num]++
	}

	muls := make([][]int, len(primes))
	active := make(map[int]bool)
	ii := make([]int, mx+1)
	for i := 0; i <= mx; i++ {
		ii[i] = -1
	}
	var it int
	for k := range freq {
		ii[k] = it
		it++
		active[k] = true
		num := k
		for num > 1 {
			x := lf[num]
			i := sort.SearchInts(primes, x)
			if len(muls[i]) == 0 {
				muls[i] = make([]int, 0, 1)
			}
			muls[i] = append(muls[i], k)
			num /= x
		}
	}

	set := NewDSU(it)

	var dfs func(num int, pos int)
	dfs = func(num int, pos int) {
		if pos == len(primes) || num*primes[pos] > mx {
			return
		}
		dfs(num, pos+1)

		var del []int
		for _, v := range muls[pos] {
			if active[v] {
				del = append(del, v)
				delete(active, v)
			}
		}
		tmp := num * primes[pos]
		if ii[tmp] >= 0 {
			for v := range active {
				set.Union(ii[tmp], ii[v])
			}
		}
		dfs(tmp, pos+1)

		for _, v := range del {
			active[v] = true
		}
	}

	dfs(1, 0)
	comps := make(map[int]int)
	var ans int
	for v := range active {
		i := ii[v]
		if set.Find(i) == i && set.cnt[i] == 1 {
			ans += freq[v]
		} else {
			comps[set.Find(i)]++
		}
	}

	return ans + len(comps)
}

func max_of(arr []int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}
