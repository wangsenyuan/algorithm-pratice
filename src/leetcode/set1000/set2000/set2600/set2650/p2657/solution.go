package p2657

func findThePrefixCommonArray(A []int, B []int) []int {

	var fa int64
	var fb int64

	var res []int

	for i := 0; i < len(A); i++ {
		fa |= 1 << (A[i] - 1)
		fb |= 1 << (B[i] - 1)
		x := fa & fb
		var cnt int
		for x > 0 {
			cnt++
			x -= x & -x
		}
		res = append(res, cnt)
	}
	return res
}
