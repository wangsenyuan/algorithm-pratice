package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	first := readNNums(reader, 5)
	n, r1, r2, r3, d := first[0], first[1], first[2], first[3], first[4]
	A := readNNums(reader, n)

	fmt.Println(solve(n, r1, r2, r3, d, A))
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

const INF = 1 << uint64(60)

func solve(n int, r1 int, r2 int, r3 int, d int, A []int) int64 {
	//dp := make([][]int64, n+1)
	//for i := 0; i <= n; i++ {
	//	dp[i] = make([]int64, 2)
	//	for j := 0; j < 2; j++ {
	//		dp[i][j] = INF
	//	}
	//}
	R1 := int64(r1)
	R2 := int64(r2)
	R3 := int64(r3)
	D := int64(d)

	a := int64(A[0])*R1 + R3
	b := min(R2, int64(A[0]+1)*R1)

	update := func(a *int64, b int64) {
		if *a > b {
			*a = b
		}
	}

	for i := 1; i < n; i++ {
		var aa, bb int64 = INF, INF
		// 0->0
		update(&aa, a+D+int64(A[i])*R1+R3)

		//0->1
		update(&bb, a+D+min(R2, int64(A[i]+1)*R1))

		//1->0
		update(&aa, b+D+R1*int64(A[i])+R3+2*D+R1)
		update(&aa, b+D+min(R1*int64(A[i]+1), R2)+D+R1+D+R1)
		//update(&dp[i+1][0], dp[i][1]+D+R2+D+R1+D+R1)

		//1->1
		update(&bb, b+D+min(R2, R1*int64(A[i]+1))+D+R1+D)
		//update(&dp[i+1][1], dp[i][1]+D++D+R1+D)

		if i == n-1 {
			update(&aa, b+D+R1*int64(A[i])+R3+D+R1)
		}
		a, b = aa, bb
	}
	// remove the first up
	return a
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
