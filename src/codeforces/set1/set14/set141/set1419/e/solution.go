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
		n := readNum(reader)
		cnt, res := solve(n)
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteString(fmt.Sprintf("\n%d\n", cnt))
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

func solve(n int) (int, []int) {
	factors := factorize(n)

	pf := make(map[int]int)

	for n%2 == 0 {
		pf[2]++
		n /= 2
	}

	for i := 3; i <= n/i; i += 2 {
		for n%i == 0 {
			pf[i]++
			n /= i
		}
	}
	if n > 1 {
		pf[n]++
	}

	pf_nums := make([]int, 0, len(pf))
	for f := range pf {
		pf_nums = append(pf_nums, f)
	}

	// 把中间的位置布置好，其他的数字可以找最小的pf放置
	// len(pf) < 30
	// 假设第一个放置p1, 最后一个放置p2, 如果有多于2个prime_factor, 则可以在0cost下进行排列
	m := len(factors)

	if len(pf_nums) == 1 {
		arr := make([]int, 0, len(factors))
		for f := range factors {
			arr = append(arr, f)
		}
		return 0, arr
	}

	if len(pf_nums) == 2 && m == 3 {
		arr := make([]int, 0, len(factors))
		for f := range factors {
			arr = append(arr, f)
		}
		// need one
		return 1, arr
	}
	sub := make([][]int, len(pf_nums))
	suf := make([]int, len(pf_nums))
	for i := 0; i < len(pf_nums); i++ {
		x := pf_nums[i]
		f := pf[x]
		sub[i] = make([]int, 0, f+1)
		for j := 0; j < f; j++ {
			sub[i] = append(sub[i], x)
			// already placed
			delete(factors, x)
			x *= pf_nums[i]
		}
		ni := (i + 1) % len(pf_nums)
		//sub[i] = append(sub[i], pf_nums[i]*pf_nums[ni])
		if ni != 0 || len(pf_nums) > 2 {
			suf[i] = pf_nums[i] * pf_nums[ni]
			delete(factors, pf_nums[i]*pf_nums[ni])
		} else {
			for a := range factors {
				if a%pf_nums[i] == 0 && a%pf_nums[ni] == 0 {
					suf[i] = a
					delete(factors, a)
					break
				}
			}
		}
	}

	for f := range factors {
		for i := 0; i < len(pf_nums); i++ {
			if f%pf_nums[i] == 0 {
				sub[i] = append(sub[i], f)
				delete(factors, f)
				break
			}
		}
	}
	ans := make([]int, 0, m)

	for i := 0; i < len(pf_nums); i++ {
		ans = append(ans, sub[i]...)
		ans = append(ans, suf[i])
	}

	return 0, ans
}

func factorize(n int) map[int]bool {
	factors := make(map[int]bool)
	for i := 2; i <= n/i; i++ {
		if n%i == 0 {
			factors[i] = true
			j := n / i
			if i != j {
				factors[j] = true
			}
		}
	}
	factors[n] = true

	return factors
}
