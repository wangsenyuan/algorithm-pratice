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
		var n int
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &n)
		var k int64
		readInt64(s, pos+1, &k)
		A := readNNums(reader, n)
		res := solve(A, k)
		buf.WriteString(fmt.Sprintf("%d\n%d %d\n", res[0], res[1], res[2]))
	}

	fmt.Print(buf.String())
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

func solve(A []int, k int64) []int {
	// gcd(A[i], A[j])
	// 做法不对
	// 比如 以3为gcd的数量中， 其实并不包括以4为gcd的数量，
	// 所以不符合二分查找
	//n := len(A)
	// 假如知道了cnt[x], 是否有助于计算cnt[y] x % y = 0, 或者 y % x = 0
	// 如果gcd(a, b) = x
	// x % y = 0, 那么 a % y, b % y = 0
	// 也就是说，如果cnt[y]表示能整除y的所有数字的对数
	ans := kthGcd(A, k)

	var arr []int
	n := len(A)
	cnt := make([]int, n+1)
	for _, num := range A {
		if num%ans == 0 {
			if cnt[num] > 1 {
				continue
			}
			cnt[num]++
			arr = append(arr, num/ans)
		}
	}
	v1, v2 := coprime(arr, len(A))
	i := findIndex(A, v1*ans)
	j := findIndex(A, v2*ans)
	if i > j {
		i, j = j, i
	}
	return []int{ans, i + 1, j + 1}
}

func findIndex(arr []int, v int) int {
	for i, num := range arr {
		if num == v {
			return i
		}
	}
	return -1
}

func coprime(arr []int, n int) (int, int) {
	spf := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if spf[i] == 0 {
			spf[i] = i
			if n/i < i {
				continue
			}
			for j := i * i; j <= n; j += i {
				if spf[j] == 0 {
					spf[j] = i
				}
			}
		}
	}
	cnt := make([]int, n+1)

	factors := make([]int, 20)
	first := -1
	for i, num := range arr {
		tmp := num
		var sz int
		for tmp > 1 {
			x := spf[tmp]
			factors[sz] = x
			sz++
			for tmp%x == 0 {
				tmp /= x
			}
		}
		var div_count int
		for state := 1; state < (1 << sz); state++ {
			prod := 1
			var bit int
			for j := 0; j < sz; j++ {
				if (state>>j)&1 == 1 {
					prod *= factors[j]
					bit++
				}
			}
			if bit&1 == 1 {
				div_count += cnt[prod]
			} else {
				div_count -= cnt[prod]
			}
			cnt[prod]++
		}
		if div_count < i {
			first = i
			break
		}
	}

	for i := 0; i < len(arr); i++ {
		if i != first && gcd(arr[first], arr[i]) == 1 {
			return arr[first], arr[i]
		}
	}
	return -1, -1
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func kthGcd(A []int, k int64) int {
	n := len(A)
	cnt := make([]int, n+1)
	// cnt[x] = 能够整除x的数字
	for i := 0; i < n; i++ {
		cnt[A[i]]++
	}

	res := make([]int64, n+1)

	for i := n; i > 0; i-- {
		var div_count int
		var rem int64
		for j := i; j <= n; j += i {
			div_count += cnt[j]
			rem += res[j]
		}
		res[i] = int64(div_count)*int64(div_count-1)/2 - rem
	}

	for i := 1; i <= n; i++ {
		if k <= res[i] {
			return i
		}
		k -= res[i]
	}

	return n
}
