package p1823

func findTheWinner(n int, k int) int {

	var loop func(n int) int

	loop = func(n int) int {
		if n == 1 {
			return 0
		}
		i := loop(n - 1)
		return (i + k) % n
	}
	res := loop(n)
	res = (res + 1) % n
	if res == 0 {
		return n
	}
	return res
}
