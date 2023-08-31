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
		a := readNNums(reader, n)
		res := solve(a)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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

func solve(a []int) bool {
	sort.Ints(a)
	// a[i] + a[j] + a[k] = a[l]
	//  考虑 a[n]
	// a[n] > 0, 如果还有两个数 i, j, a[i] + a[j] > 0, 则不可能，
	// a[0] + a[1] + a[2] 如果在a中
	// 假设在2的后
	// a[i] = a[0] + a[1] + a[2]
	//      a[i] + a[1] + a[2] => a[0] + 2 * (a[1] + a[2])
	// a[i] >= 0, 因为如果 a[i] < 0, 那么 0, 1, 2都< 0, 但是，我们直到0是最小的， 三个负数相加只会更小
	// a[i] > 0, 使用 a[i] 替代a[0], => 更快的得到正数， 进而得到更多的正数
	// 如果i在2的后面，必然是不行的
	n := len(a)
	if a[2] < 0 {
		// three negative
		return false
	}
	if a[n-3] > 0 {
		// three positive
		return false
	}
	var arr []int

	var cnt int
	for i := 0; i < n; i++ {
		if a[i] == 0 {
			cnt++
		} else {
			arr = append(arr, a[i])
		}
	}
	if cnt > 2 {
		cnt = 2
	}

	for i := 0; i < cnt; i++ {
		arr = append(arr, 0)
	}
	// len(arr) <= 6
	set := make(map[int]bool)

	for _, x := range arr {
		set[x] = true
	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				sum := arr[i] + arr[j] + arr[k]
				if !set[sum] {
					return false
				}
			}
		}
	}

	return true
}
