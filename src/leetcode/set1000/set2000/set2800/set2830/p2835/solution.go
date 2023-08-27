package p2835

const H = 31

func minOperations(nums []int, target int) int {
	freq := make([]int, H+1)

	for _, num := range nums {
		var d int
		for 1<<d < num {
			d++
		}
		// 1 << d = num
		freq[d]++
	}
	var ans int
	for i := 0; i < H; i++ {
		if (target>>i)&1 == 1 {
			j := i
			for j < H && freq[j] == 0 {
				j++
			}
			if j == H {
				return -1
			}
			if i < j {
				freq[j]--
				for k := i + 1; k < j; k++ {
					freq[k]++
				}
				freq[i] = 2
				ans += j - i
			}
			freq[i]--
		}
		freq[i+1] += freq[i] / 2
	}

	return ans
}
