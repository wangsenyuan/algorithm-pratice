package p986

/**
 * Definition for an interval.
 */
type Interval struct {
	Start int
	End   int
}

func intervalIntersection(A []Interval, B []Interval) []Interval {
	var i, j int

	res := make([]Interval, 0, len(A))
	for i < len(A) && j < len(B) {
		a := A[i]
		for j < len(B) && B[j].End < a.Start {
			//B[j] is before a
			j++
		}
		if j == len(B) {
			break
		}
		//B[j].End >= a.Start
		tmp := intersect(a, B[j])
		if tmp.Start <= tmp.End {
			res = append(res, intersect(a, B[j]))
		}

		if B[j].End >= a.End {
			i++
		} else {
			j++
		}
	}

	return res
}

func iv(a, b int) Interval {
	return Interval{a, b}
}

func intersect(a, b Interval) Interval {
	start := max(a.Start, b.Start)
	end := min(a.End, b.End)
	return iv(start, end)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
