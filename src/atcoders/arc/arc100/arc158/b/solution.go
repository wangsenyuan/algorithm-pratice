package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)

	res := solve(A)
	fmt.Printf("%.13f\n", res[0])
	fmt.Printf("%.13f\n", res[1])
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(x []int) []float64 {
	sort.Ints(x)

	n := len(x)
	var arr []int
	for i := 0; i <= n; i++ {
		if x[i] > 0 {
			// [0...i) are negative
			if i < 7 {
				arr = append(arr, x[:i]...)
			} else {
				arr = append(arr, x[:3]...)
				arr = append(arr, x[i-3:i]...)
			}

			if n-i < 7 {
				arr = append(arr, x[i:]...)
			} else {
				arr = append(arr, x[i:i+3]...)
				arr = append(arr, x[n-3:]...)
			}
			break
		}
	}
	if len(arr) == 0 {
		// all negative
		if n < 7 {
			arr = x
		} else {
			arr = append(arr, x[:3]...)
			arr = append(arr, x[n-3:]...)
		}
	}

	// len(arr) <= 12
	ans := []float64{math.MaxFloat64, math.MinInt64}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				tmp := float64(arr[i]+arr[j]+arr[k]) / (float64(arr[i]) * float64(arr[j]) * float64(arr[k]))
				ans[0] = math.Min(ans[0], tmp)
				ans[1] = math.Max(ans[1], tmp)
			}
		}
	}

	return ans
}
