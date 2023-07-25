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

	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)

		A := readNNums(reader, n)

		res := solve(k, A)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(k int, bonus []int) int64 {

	// 贪心算法是最麻烦的，贪心的某些方面在部分情况下是有效的，但是在其他情况下失效
	// 观察1，如果不存在负值（或者负值的个数<=k)，那么完全不需要reset (先reset掉全部的负值）
	// 观察2，如果全部是负值，就需要全部reset，且分成k+1个组，在每个组内的数字大体接近
	// 观察3，处在中间状态时，这时都不清楚是否应该将全部reset掉？
	//   这是因为 考虑只由[1,-1]组成的长度n数组 只要sum >= 0, 就不需要reset （事实上也是需要reset的，reset一次，相当于相当于少了一个-1）
	// 观察4，假设将数组分成了x组, (x <= k + 1), 那么在每一组中，始终先处理大的数，再处理小的数是更优的
	// 观察5，假设有两个组，一个组有a个正数组成，另外一个组有b个负值组成
	// [3, 1, x], [-1, -2, y]
	// 这时它们的结果是 2 * 3 + 1 + -2 * 2 + -2 = 1
	// 如果把它们合并后呢？
	// 3 * 5 + 1 * 4 + -1 * 3 + -2 * 2 + max(x, y)
	// = 12 + max(x, y)
	// 也就是说 只要 max(x, y) >= -11, 合并在一个数组内就是更优的结果
	// 当然这里还需要知道 连续负值的贡献和正值的贡献
	// 先把所有的正值+最小的负值放在一个group中 （为啥要这么做?)
	// 然后尝试向里面放置负值 (选择哪个负值 ?)
	//  选择大的负值更好， 只要这个abs(x) <= sum(pref)
	// 但是为啥要把正值放在一个group中呢？ 正值放在一起，在后面添加负值的时候，才能增大正值的作用。
	sort.Ints(bonus)
	reverse(bonus)

	if k == 0 {
		return solve0(bonus)
	}

	pq := make(IntHeap, k+1)
	heap.Init(&pq)
	n := len(bonus)

	var res int64

	for i := 0; i < n; i++ {
		x := heap.Pop(&pq).(int64)
		res += x
		heap.Push(&pq, x+int64(bonus[i]))
	}

	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func solve0(bonus []int) int64 {

	var sum int64
	var ans int64
	n := len(bonus)
	for i := 0; i < n; i++ {
		ans += sum
		sum += int64(bonus[i])
	}
	return ans
}

// An IntHeap is a min-heap of ints.
type IntHeap []int64

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int64))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
