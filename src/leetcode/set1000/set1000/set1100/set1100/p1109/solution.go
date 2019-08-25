package p1109

func corpFlightBookings(bookings [][]int, n int) []int {
	sum := make([]int, n+1)
	for _, booking := range bookings {
		i, j, k := booking[0]-1, booking[1]-1, booking[2]
		sum[i] += k
		sum[j+1] -= k
	}

	for i := 1; i < n; i++ {
		sum[i] += sum[i-1]
	}
	return sum[:n]
}
