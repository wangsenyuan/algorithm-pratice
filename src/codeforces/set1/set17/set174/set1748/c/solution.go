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
		a := readNNums(reader, n)
		res := solve(a)
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

func solve(v []int) int {
	// 这个太精巧了
	// 考虑 s1 = a1
	//     s2 = a1 + a2
	//     si = a1 + a2 + ... + ai
	//  对于si组成的新的数组 [s1, s2, ... si, ... sn]
	// 可以按照ax = 0 进行分组
	// 假设分成了组
	//  [s1, s2, ... si-1]
	//  [si, ....  sj]
	//  ...
	// 这里 a[i] = 0, 之所以要按照0分组，是因为，可以通过改变a[i] 使得它同组的分段sum变成0
	// 假设 a[i+1] + a[i+2] + .. a[i+k] = x, 那么通过将 a[i] = -x
	// 就可以得到 a[i] + ... a[i+k] = 0
	// 此时最优的-x是多少呢？就是该分组内出现最多的数，此时，就可以得到增加出现次数的分组和0
	freq := make(map[int64]int)
	var mf int
	var sum int64
	var found0 bool
	var ans int
	for _, a := range v {
		if a == 0 {
			if found0 {
				ans += mf
			} else {
				ans += freq[0]
			}
			found0 = true
			mf = 0
			freq = make(map[int64]int)
		}
		sum += int64(a)
		freq[sum]++
		if freq[sum] > mf {
			mf = freq[sum]
		}
	}
	if found0 {
		ans += mf
	} else {
		ans += freq[0]
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
