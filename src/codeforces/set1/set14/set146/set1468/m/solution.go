package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
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
		sets := make([][]int, n)
		for i := 0; i < n; i++ {
			s, _ := reader.ReadBytes('\n')
			var m int
			pos := readInt(s, 0, &m)
			sets[i] = make([]int, m)
			for j := 0; j < m; j++ {
				pos = readInt(s, pos+1, &sets[i][j])
			}
		}
		res := solve(sets)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d %d\n", res[0], res[1]))
		}
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

func solve(sets [][]int) []int {
	nums := make(map[int]int)

	for _, set := range sets {
		for _, x := range set {
			nums[x]++
		}
	}
	m := len(nums)
	d := int(math.Sqrt(float64(m)))

	arr := make([]int, 0, len(nums))
	for num := range nums {
		arr = append(arr, num)
	}
	for i, num := range arr {
		nums[num] = i
	}

	var large []int
	var small []int

	n := len(sets)

	for i := 0; i < n; i++ {
		for j := 0; j < len(sets[i]); j++ {
			sets[i][j] = nums[sets[i][j]]
		}
		if len(sets[i]) >= d {
			large = append(large, i)
		} else {
			small = append(small, i)
		}
	}
	w := make([]int, m)
	for _, i := range large {
		for _, x := range sets[i] {
			w[x]++
		}
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			var cnt int
			for k := 0; k < len(sets[j]) && cnt < 2; k++ {
				cnt += w[sets[j][k]]
			}
			if cnt == 2 {
				return []int{i + 1, j + 1}
			}
		}
		for _, x := range sets[i] {
			w[x]--
		}
	}

	type Pair struct {
		first  int
		second int
	}

	fx := make([][]Pair, m)

	for _, i := range small {
		sort.Ints(sets[i])
		for j := 0; j < len(sets[i]); j++ {
			x := sets[i][j]
			for k := j + 1; k < len(sets[i]); k++ {
				y := sets[i][k]
				if len(fx[x]) == 0 {
					fx[x] = make([]Pair, 0, 1)
				}
				fx[x] = append(fx[x], Pair{y, i})
			}
		}
	}

	for i := 0; i < m; i++ {
		sort.Slice(fx[i], func(a, b int) bool {
			return fx[i][a].first < fx[i][b].first
		})

		for j := 1; j < len(fx[i]); j++ {
			if fx[i][j].first == fx[i][j-1].first {
				return []int{fx[i][j].second + 1, fx[i][j-1].second + 1}
			}
		}
	}

	return nil
}
