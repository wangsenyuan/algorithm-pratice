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
		n := readNInt64s(reader, 1)[0]
		res := solve(n)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

const N = 1_00_100

var primes []int

func init() {
	set := make([]bool, N)
	for x := 2; x < N; x++ {
		if !set[x] {
			primes = append(primes, x)
			if N/x <= x {
				continue
			}

			for j := x * x; j < N; j += x {
				set[j] = true
			}
		}
	}
}

func solve(n int64) []int64 {
	// 如果是质数，那么就是n
	// a1 * a2 * .. * ak = n
	// a2 % a1 = 0
	// a3 % a2 = 0
	// n % ((a1) ** k) = 0
	// a ** k1 *  b ** k2  * c ** k3 ... = n
	// 还是要从最小的factor开始

	tmp := n

	factors := make(map[int64]int)

	for i := 0; i < len(primes) && int64(primes[i]) <= tmp; i++ {
		x := int64(primes[i])
		for tmp%x == 0 {
			factors[x]++
			tmp /= x
		}
	}

	if tmp > 1 {
		factors[tmp]++
	}
	ps := make([]Pair, 0, len(factors))

	for k, v := range factors {
		ps = append(ps, Pair{k, v})
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].second > ps[j].second
	})

	k := ps[0].second
	res := make([]int64, k)
	for i := 0; i < k; i++ {
		res[i] = ps[0].first
		n /= res[i]
	}

	res[k-1] *= n

	return res
}

type Pair struct {
	first  int64
	second int
}
