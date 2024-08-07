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
		p := readNNums(reader, n)
		res := solve(p)
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

func solve(p []int) int {
	n := len(p)
	v1 := make([][]int, n+1)
	v2 := make([][]int, n+1)
	cnt := make([]int, n+1)
	has := make([]int, n+1)

	for i := 1; i <= n; i++ {
		g := gcd(i, p[i-1])
		v1[i/g] = append(v1[i/g], p[i-1]/g)
		v2[p[i-1]/g] = append(v2[p[i-1]/g], i/g)
	}

	ans := ^len(v1[1]) + 1

	save := func(d int, i int) {
		if has[d] != i {
			has[d] = i
			cnt[d] = 0
		}
		cnt[d]++
	}

	for i := 1; i <= n; i++ {
		for _, j := range v1[i] {
			for k := 1; k*k <= j; k++ {
				if j%k == 0 {
					save(k, i)
					if k*k != j {
						save(j/k, i)
					}
				}
			}
		}
		for j := i; j <= n; j += i {
			for _, k := range v2[j] {
				if has[k] == i {
					ans += cnt[k]
				}
			}
		}
	}

	return ans / 2
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
