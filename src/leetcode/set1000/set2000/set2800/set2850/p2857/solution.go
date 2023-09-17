package p2857

func countPairs(coordinates [][]int, k int) int {

	freq := make(map[Pair]int)
	var res int

	for _, coord := range coordinates {
		x, y := coord[0], coord[1]
		for i := 0; i <= k; i++ {
			// x ^ x1 = i
			x1 := x ^ i
			y1 := y ^ (k - i)
			res += freq[Pair{x1, y1}]
		}
		freq[Pair{x, y}]++
	}

	return res
}

type Pair struct {
	first  int
	second int
}
