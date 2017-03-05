package main

import "fmt"

func main() {
	picture := [][]byte{
		[]byte("WBBWWBWWWWWBBWW"),
		[]byte("WBBWWBWWWWWBBWW"),
		[]byte("WWWWWBBBWBWWWWB"),
		[]byte("WWBWBWWWWBBWBWW"),
		[]byte("WBBWWBWWWWWBBWW"),
		[]byte("WWBWBWWWWBBWBWW"),
		[]byte("WWBWBWWWWBBWBWW"),
		[]byte("WWBWBWWWWBBWBWW"),
	}
	fmt.Println(findBlackPixel(picture, 5))
}

func findBlackPixel(picture [][]byte, N int) int {
	n := len(picture)
	if n == 0 {
		return 0
	}

	m := len(picture[0])
	if m == 0 {
		return 0
	}

	row := make([]int, n)
	col := make([]int, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if picture[i][j] == 'B' {
				row[i]++
				col[j]++
			}
		}
	}

	var res int

	for j := 0; j < m; j++ {
		if col[j] != N {
			continue
		}

		a := -1
		for i := 0; i < n; i++ {
			if row[i] == N && picture[i][j] == 'B' {
				a = i
				break
			}
		}
		if a == -1 {
			continue
		}

		as := string(picture[a])
		same := true
		for i := 0; i < n && same; i++ {
			if picture[i][j] == 'B' {
				same = string(picture[i]) == as
			}
		}
		if same {
			//fmt.Printf("%d %d %s\n", j, a, as)
			res += N
		}
	}

	return res
}
