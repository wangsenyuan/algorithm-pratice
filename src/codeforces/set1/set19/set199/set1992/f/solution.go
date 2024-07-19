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
		n, x := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(a, x)
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

func solve(a []int, x int) int {

	var res int
	var divs []int
	divs = append(divs, 1)

	for _, num := range a {
		if num == 1 || x%num != 0 {
			continue
		}
		j := sort.SearchInts(divs, x/num)
		if j < len(divs) && divs[j] == x/num {
			res++
			divs = divs[:0]
			divs = append(divs, 1, num)
		} else {
			p := len(divs)
			for i := 0; i < p; i++ {
				divs = append(divs, divs[i]*num)
			}
			divs = sortAndDistinct(divs)
		}
	}
	return res + 1
}

func sortAndDistinct(arr []int) []int {
	sort.Ints(arr)
	n := len(arr)
	var m int
	for i := 1; i <= n; i++ {
		if i == n || arr[i] > arr[i-1] {
			arr[m] = arr[i-1]
			m++
		}
	}
	return arr[:m]
}
