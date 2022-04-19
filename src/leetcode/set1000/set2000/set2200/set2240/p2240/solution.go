package p2240

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {

	if cost1 < cost2 {
		cost1, cost2 = cost2, cost1
	}

	var res int64

	for x := 0; x <= total/cost1; x++ {
		y := (total - x*cost1) / cost2
		res += int64(y + 1)
	}

	return res
}
