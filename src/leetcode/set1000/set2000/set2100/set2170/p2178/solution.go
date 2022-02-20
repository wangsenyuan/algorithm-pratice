package p2178

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 == 1 {
		return nil
	}
	// res would be [2, 4, 6, 8, ....]
	// n would be roughly sqrt of finalSum
	// n = finalSum / 2
	// then split n into distinct size of groups
	var res []int64
	var i int64 = 1
	var sum int64
	for {
		tmp := finalSum - sum
		if tmp < i*2+(i+1)*2 {
			res = append(res, tmp)
			break
		}
		res = append(res, i*2)
		sum += i * 2
		i++
	}
	return res
}
