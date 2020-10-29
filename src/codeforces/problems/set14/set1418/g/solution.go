package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	A := readNNums(reader, n)

	fmt.Println(solve(n, A))
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

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(99))
}

func solve(n int, A []int) int64 {
	hashes := make([]uint64, n+1)
	indices := make([][]int, n+1)
	freq := make([]uint64, n+1)
	hashValues := make([]uint64, n+1)
	hashFreq := make(map[uint64]int)
	for i := 1; i <= n; i++ {
		hashValues[i] = r.Uint64()
		indices[i] = make([]int, 0, 4)
	}
	hashFreq[0]++
	var start int
	var ans int64
	for i := 0; i < n; i++ {
		cur := A[i]
		if len(indices[cur]) >= 3 {
			remove := indices[cur][0]
			indices[cur] = indices[cur][1:]
			for start <= remove {
				hashFreq[hashes[start]]--
				start++
			}
		}
		before := freq[cur]
		freq[cur] = (before + 1) % 3
		hashes[i+1] = hashes[i] + (freq[cur]-before)*hashValues[cur]
		ans += int64(hashFreq[hashes[i+1]])
		hashFreq[hashes[i+1]]++
		indices[cur] = append(indices[cur], i)
	}

	return ans
}
