package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(m, a)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const mod = 998244353

func mul(a, b int) int {
	return int(int64(a) * int64(b) % mod)
}

func solve(m int, a []int) int {
	ps := getPrimeFactors(a[0])

	n := len(a)

	res := 1

	for i := 1; i < n; i++ {
		if a[i-1]%a[i] != 0 {
			return 0
		}
		left := a[i-1] / a[i]

		if left == 1 {
			res = mul(res, m/a[i])
			continue
		}

		till := m / a[i]

		var nums []int
		for v := range ps {
			if left%v == 0 {
				nums = append(nums, v)
			}
		}
		sz := len(nums)

		var ans int
		for state := 0; state < (1 << sz); state++ {
			prod := 1
			var cnt int
			for j := 0; j < sz; j++ {
				if (state>>j)&1 == 1 {
					cnt++
					prod *= nums[j]
				}
			}
			if cnt&1 == 0 {
				ans += till / prod
			} else {
				ans -= till / prod
			}
		}
		res = mul(res, ans)
	}

	return res
}

type Pair struct {
	first  int
	second int
}

func getPrimeFactors(num int) map[int]int {
	res := make(map[int]int)

	for i := 2; i*i <= num; i++ {
		var cnt int
		for num%i == 0 {
			cnt++
			num /= i
		}
		if cnt > 0 {
			res[i] = cnt
		}
	}
	if num > 1 {
		res[num] = 1
	}
	return res
}
