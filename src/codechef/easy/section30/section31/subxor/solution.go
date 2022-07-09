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
		s := readString(reader)
		res := solve(n, s)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const MOD = 998244353

func solve(n int, s string) int {

	flip := make([]int, n+2)

	// 考虑 10， 考虑1
	// 1的贡献为 (1 << 2) - 1, = 3
	// 1111   s(0) => 1111
	//        s(1) =>  111
	//        s(2) =>   11
	//        s(3) =>    1

	for i := 0; i < n; i++ {
		if s[i] == '1' {
			flip[n-i-1] += i + 1
		}
	}
	var res int
	for i := n - 1; i >= 0; i-- {
		flip[i] += flip[i+1]
		if flip[i]%2 == 1 {
			// contributes 1 << i
			res = modAdd(res, pow(2, i))
		}
	}

	return res
}

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	r %= MOD
	return int(r)
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = modMul(res, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return res
}

func samples() {
	for x := 0; x < 1000; x++ {
		sample(x)
	}
}

func sample(x int) int {
	arr := digitsArray(x)
	var res int
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			var num int
			for k := i; k <= j; k++ {
				num = num*2 + arr[k]
			}
			res ^= num
		}
	}

	// rra := digitsArray(res)

	// fmt.Printf("%s -> %s\n", toString(arr), toString(rra))
	return res
}

func digitsArray(num int) []int {
	arr := make([]int, 10)
	for i := 9; i >= 0; i-- {
		arr[i] = (num >> uint(9-i)) & 1
	}
	return arr
}

func toString(arr []int) string {
	var buf bytes.Buffer

	for i := 0; i < len(arr); i++ {
		buf.WriteByte(byte('0' + arr[i]))
	}

	return buf.String()
}
