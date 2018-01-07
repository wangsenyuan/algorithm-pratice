package p759

type Interval struct {
	Start int
	End   int
}

func employeeFreeTime(avails [][]Interval) []Interval {

	merge := func(a, b []Interval) []Interval {
		res := make([]Interval, 0, len(a)+len(b))

		i, j := 0, 0
		begin, end := -1, -1
		for i < len(a) && j < len(b) {
			x, y := a[i], b[j]
			var cur Interval
			if overlap(x, y) {
				tmp1 := min(x.Start, y.Start)
				tmp2 := max(x.End, y.End)
				cur = Interval{tmp1, tmp2}
				i++
				j++
			} else if x.End < y.Start {
				cur = x
				i++
			} else {
				cur = y
				j++
			}

			if cur.Start <= end {
				//extend
				end = max(cur.End, end)
			} else {
				if begin >= 0 {
					res = append(res, Interval{begin, end})
				}
				begin = cur.Start
				end = cur.End
			}
		}

		for i < len(a) {
			cur := a[i]
			if cur.Start <= end {
				end = max(cur.End, end)
			} else {
				if begin >= 0 {
					res = append(res, Interval{begin, end})
				}
				begin = cur.Start
				end = cur.End
			}
			i++
		}

		for j < len(b) {
			cur := b[j]
			if cur.Start <= end {
				end = max(cur.End, end)
			} else {
				if begin >= 0 {
					res = append(res, Interval{begin, end})
				}
				begin = cur.Start
				end = cur.End
			}
			j++
		}

		if begin > 0 {
			res = append(res, Interval{begin, end})
		}

		return res
	}

	var fn func(avalis [][]Interval) []Interval

	fn = func(avails [][]Interval) []Interval {
		n := len(avails)
		if n == 1 {
			return avails[0]
		}

		mid := n / 2
		a := fn(avails[:mid])
		b := fn(avails[mid:])
		return merge(a, b)
	}

	xs := fn(avails)
	ys := make([]Interval, 0, len(xs))

	for i := 1; i < len(xs); i++ {
		a, b := xs[i-1], xs[i]
		ys = append(ys, Interval{a.End, b.Start})
	}

	return ys
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func overlap(a, b Interval) bool {
	return (a.Start <= b.Start && b.Start <= a.End) ||
		(b.Start <= a.Start && a.Start <= b.End)
}
