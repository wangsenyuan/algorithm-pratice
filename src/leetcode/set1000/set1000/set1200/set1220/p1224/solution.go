package p1224

func maxEqualFreq(nums []int) int {
	var ans int

	n := len(nums)

	freq := make(map[int]int)
	cnt := make(map[int]int)

	for i := 0; i < n; i++ {
		old := cnt[nums[i]]

		if v, found := freq[old]; found {
			if v > 1 {
				freq[old] = v - 1
			} else {
				delete(freq, old)
			}
		}

		cnt[nums[i]]++

		freq[cnt[nums[i]]]++

		if len(freq) == 2 {
			// a candidate found
			var first, second int
			for a := range freq {
				if first == 0 {
					first = a
				} else {
					second = a
				}
			}

			if freq[first] == 1 && first == 1 || freq[second] == 1 && second == 1 {
				// can remove that
				ans = i + 1
				continue
			}
			if first-second == 1 && freq[first] == 1 {
				ans = i + 1
				continue
			}

			if second-first == 1 && freq[second] == 1 {
				ans = i + 1
				continue
			}
		} else if len(freq) == 1 {
			// special case
			var first int
			for a := range freq {
				first = a
			}
			if first == 1 || freq[first] == 1 {
				// can remove any number
				ans = i + 1
			}
		}

	}

	return ans
}
