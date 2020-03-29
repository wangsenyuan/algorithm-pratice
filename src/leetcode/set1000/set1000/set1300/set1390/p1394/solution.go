package p1394

func numTeams(rating []int) int {
	n := len(rating)

	var res int

	for i := 1; i < n-1; i++ {
		var cnt1, cnt2 int
		for j := 0; j < i; j++ {
			if rating[j] < rating[i] {
				cnt1++
			} else if rating[j] > rating[i] {
				cnt2++
			}
		}

		var cnt3, cnt4 int

		for j := n - 1; j > i; j-- {
			if rating[j] > rating[i] {
				cnt3++
			} else if rating[j] < rating[i] {
				cnt4++
			}
		}

		res += cnt1*cnt3 + cnt2*cnt4
	}

	return res
}

type BIT struct {
	arr  []int
	size int
}

func CreateBIT(n int) BIT {
	arr := make([]int, n+1)
	return BIT{arr, n}
}

func (bit *BIT) Update(pos int, val int) {
	if pos == 0 {
		return
	}
	arr := bit.arr
	size := bit.size
	for pos <= size {
		arr[pos] += val
		pos += pos & (-pos)
	}
}

func (bit *BIT) Get(pos int) int {
	arr := bit.arr
	var res int
	if pos > bit.size {
		return 0
	}
	for pos > 0 {
		res += arr[pos]
		pos -= pos & (-pos)
	}
	return res
}
