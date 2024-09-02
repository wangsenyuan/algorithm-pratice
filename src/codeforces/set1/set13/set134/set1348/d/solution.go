package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		res := solve(n)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func solve(n int) []int {
	var arr []int
	arr = append(arr, 1)
	sum := 1
	val := 1
	for sum < n {
		val *= 2
		if sum+val <= n {
			sum += val
			arr = append(arr, val)
		} else {
			arr = append(arr, n-sum)
			sum = n
		}
	}

	sort.Ints(arr)
	res := make([]int, len(arr)-1)
	for i := 0; i+1 < len(arr); i++ {
		res[i] = arr[i+1] - arr[i]
	}
	return res
}

func solve1(n int) []int {

	check := func() int {
		cnt := 1
		tot := 1
		for i := 0; ; i++ {
			cnt *= 2
			tot += cnt
			if tot >= n {
				return i + 1
			}
		}
	}

	x := check()

	// 最少x天，可以完成
	var res []int
	cnt := 1
	tot := 1

	for i := 0; i < x; i++ {
		// 在不分裂的情况下，接下来将增加这么多
		tmp := cnt * (x - i)
		// 在全部分裂的情况下，将增加这么多
		tmp2 := 2 * tmp
		if tot+tmp2 <= n {
			res = append(res, cnt)
			cnt *= 2
			tot += cnt
			continue
		}
		// tot + tmp2 > n
		diff := n - (tot + tmp)
		tmp3 := diff / (x - i)
		res = append(res, tmp3)
		cnt += tmp3
		tot += cnt
	}

	return res
}
