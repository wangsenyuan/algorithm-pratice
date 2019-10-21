package p1227

func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1.0
	}
	// if first person take the right one, 1 / n
	// 有 m 个人坐错了位置
	// .....

	// 第一个人有 1/n 的概率take seat n, else (n - 2) / (n - 1) 的概率坐错
	// (n - 2) / n * 1 / (n - 1) take seat n else
	//
	return 0.5
}
