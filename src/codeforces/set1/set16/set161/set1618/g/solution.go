package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m, k := readThreeNums(reader)
	A := readNNums(reader, n)
	B := readNNums(reader, m)
	K := readNNums(reader, k)
	res := solve(A, B, K)

	var buf bytes.Buffer

	for _, num := range res {
		buf.WriteString(fmt.Sprintf("%d\n", num))
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
func solve(A []int, B []int, K []int) []int64 {
	sort.Ints(A)
	sort.Ints(B)

	//假设alice，一开始有x，交换某个y1, y1 <= x + k
	// 如果仍然有y2 <= y1 + k, 可以用y1交换y2回来
	// 所以，如果有一段 [yi...yj], 都满足 yi <= yi-1 + k
	// 那么使用某个x  x >= yi - k， 可以把yj交换回来
	// 同样的如果有 xi...xj,如果把xj交换过去了，最后也可以用xi交换回来
	// 对于给定的k， 把x和y合并在一起后，相邻数字间的差 <= k, 分组，从各分组中，取出其中最大的x分组大小的数
	// 可以在O(n)时间内解决，
	//但是如何处理所有的k呢？
	// 还是考虑分组，当k越来越小的时候， 分组会分开，（越来越大的时候， 分组会合并）
	// 考虑k越来越大， 分组合并的情况， 假设两个分组合并了，那么先减去它们的贡献，然后加上这个分组内新的值
	// 最多合并n-1次， 只要能快速进行查找（哪些需要合并), 合并，就可以了
	// 然后就是如何快速的计算，分组内有多少个x元素，这个可以用bit计算
	// 还需要计算分组内的最大的x个数字的和， 用seg-tree？
	kp := make([]Pair, len(K))
	for i := 0; i < len(K); i++ {
		kp[i] = Pair{K[i], i}
	}

	sort.Sort(Pairs(kp))

	nums := make([]Pair, len(A)+len(B))
	n := len(A)

	for i := 0; i < n; i++ {
		nums[i] = Pair{A[i], i}
	}

	for i, j := n, 0; j < len(B); i, j = i+1, j+1 {
		nums[i] = Pair{B[j], i}
	}
	sort.Sort(Pairs(nums))
	// when second < n, it is x
	pq := make(PQ, len(nums)-1)
	for i := 1; i < len(nums); i++ {
		pq[i-1] = Item{i, nums[i].first - nums[i-1].first}
	}

	m := len(nums)

	heap.Init(&pq)

	var sum int64

	set := make([]int, m)
	next := make([]int, m)

	cnt := make([]int, m)
	res := make([]int64, m)
	pref := make([]int64, m)
	for i, num := range nums {
		if num.second < n {
			cnt[i] = 1
			res[i] = int64(num.first)
		}
		set[i] = i
		next[i] = i
		sum += res[i]
		pref[i] += int64(num.first)
		if i > 0 {
			pref[i] += pref[i-1]
		}
	}

	var find func(arr []int, i int) int

	find = func(arr []int, i int) int {
		if arr[i] != i {
			arr[i] = find(arr, arr[i])
		}
		return arr[i]
	}

	ans := make([]int64, len(K))

	for _, cur := range kp {
		k := cur.first
		for pq.Len() > 0 && pq[0].dist <= k {
			tp := heap.Pop(&pq).(Item)
			// 需要合并
			i := tp.id
			// connect id & id-1
			j := find(set, i-1)
			x := find(next, i)
			cnt[j] += cnt[i]
			set[i] = i - 1
			next[i-1] = i

			sum -= res[i]
			sum -= res[j]

			res[j] = pref[x]

			if x-cnt[j] >= 0 {
				res[j] -= pref[x-cnt[j]]
			}

			sum += res[j]
		}
		ans[cur.second] = sum
	}

	return ans
}

type Pair struct {
	first  int
	second int
}

func (this Pair) Less(that Pair) bool {
	return this.first < that.first || this.first == that.first && this.second < that.second
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].Less(ps[j])
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type Item struct {
	id   int
	dist int
}

type PQ []Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	i := x.(Item)
	*pq = append(*pq, i)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	rt := old[n-1]
	*pq = old[:n-1]
	return rt
}
