package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		ops := make([][]int, m)
		for i := 0; i < m; i++ {
			s, _ := reader.ReadBytes('\n')
			if s[0] == '1' {
				var x int
				readInt(s, 2, &x)
				ops[i] = []int{1, x}
			} else {
				ops[i] = []int{2}
			}
		}
		res := solve(n, ops)
		for _, x := range res {
			if x {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
		}
	}

	fmt.Print(buf.String())
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
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const N = 1000010

var lpf [N]int

func init() {
	var primes []int
	for i := 2; i < N; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}

		for j := 0; j < len(primes); j++ {
			if i*primes[j] >= N {
				break
			}
			lpf[i*primes[j]] = primes[j]
			if i%primes[j] == 0 {
				break
			}
		}
	}
}

func solve(n int, ops [][]int) []bool {
	cnt := 1
	var fs map[int]int

	reset := func() {
		fs = getFactors(n)
		cnt = 1
		for _, v := range fs {
			cnt *= v + 1
		}
	}

	reset()

	var res []bool

	for _, cur := range ops {
		if cur[0] == 1 {
			x := cur[1]
			xs := getFactors(x)

			ok := true

			for k := range xs {
				cnt /= fs[k] + 1
			}

			for k, v := range xs {
				fs[k] += v
				cnt *= fs[k] + 1
			}
			// cnt 现在是n的因子的数量
			// 这里希望cnt是n的一个因子
			ys := getFactors2(cnt)
			for k, v := range ys {
				if fs[k] < v {
					ok = false
					break
				}
			}
			res = append(res, ok)
		} else {
			reset()
		}
	}

	return res
}

func getFactors2(num int) map[int]int {
	if num < N {
		return getFactors(num)
	}
	res := make(map[int]int)
	for x := 2; x <= num/x; x++ {
		for num%x == 0 {
			res[x]++
			num /= x
		}
	}
	if num > 1 {
		res[num]++
	}
	return res
}

func getFactors(n int) map[int]int {
	res := make(map[int]int)
	for n > 1 {
		res[lpf[n]]++
		n /= lpf[n]
	}
	return res
}
