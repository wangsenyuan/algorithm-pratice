package p2845

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	freq := make(map[int]int)

	var cnt int

	var ans int64

	freq[0]++

	for _, num := range nums {
		x := num % modulo
		if x == k {
			cnt++
		}

		a := cnt % modulo

		// want to find another prefix hav (cnt - count(prefix)) % module = k
		ans += int64(freq[(a-k+modulo)%modulo])

		freq[a]++
	}

	return ans
}
