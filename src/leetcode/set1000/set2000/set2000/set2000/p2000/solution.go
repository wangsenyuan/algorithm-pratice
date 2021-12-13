package p2000

func interchangeableRectangles(rectangles [][]int) int64 {

	cnt := make(map[Pair]int)
	n := len(rectangles)
	for i := 0; i < n; i++ {
		cur := rectangles[i]
		w, h := cur[0], cur[1]
		c := gcd(w, h)
		p := Pair{w / c, h / c}
		cnt[p]++
	}

	var res int64

	for _, v := range cnt {
		res += int64(v) * int64(v-1) / 2
	}

	return res
}

type Pair struct {
	first, second int
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
