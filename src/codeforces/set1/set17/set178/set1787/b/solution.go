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
		n := readNum(reader)
		res := solve(n)
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

func solve(n int) int {
	// n的质数化，
	// 100 = 2 * 2 * 5 * 5
	// 有好几种选择 2 + 2 + 5 + 5 = 14
	// 2 * 5 + 2 * 5 = 20
	// 2 * 5 * 2 = 20
	// 显然p值越大越好
	// 然后相同的值越大越好
	factors := make(map[int]int)

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				factors[i]++
				n /= i
			}
		}
	}
	if n > 1 {
		factors[n]++
	}

	arr := make([]Pair, 0, len(factors))
	prod := 1
	for k, v := range factors {
		prod *= k
		arr = append(arr, Pair{k, v})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].second < arr[j].second
	})

	var ans int
	var prev int
	for i := 0; i < len(arr); {
		ans += prod * (arr[i].second - prev)
		j := i
		for i < len(arr) && arr[j].second == arr[i].second {
			prod /= arr[i].first
			i++
		}
		prev = arr[j].second
	}

	return ans
}

type Pair struct {
	first  int
	second int
}
