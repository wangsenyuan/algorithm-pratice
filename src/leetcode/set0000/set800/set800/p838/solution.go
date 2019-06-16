package p838

import "bytes"

func pushDominoes(dominoes string) string {
	n := len(dominoes)
	L, R := make([]int, n), make([]int, n)

	INF := 1 << 25
	if dominoes[0] == 'R' {
		R[0] = 0;
	} else {
		// far far away
		R[0] = INF;
	}

	for i := 1; i < n; i++ {
		if dominoes[i] == '.' {
			R[i] = R[i-1] + 1
		} else if dominoes[i] == 'R' {
			R[i] = 0
		} else {
			// push L
			R[i] = INF
		}
	}
	if dominoes[n-1] == 'L' {
		L[n-1] = 0
	} else {
		L[n-1] = INF
	}

	for i := n - 2; i >= 0; i-- {
		if dominoes[i] == '.' {
			L[i] = L[i+1] + 1
		} else if dominoes[i] == 'L' {
			L[i] = 0
		} else {
			L[i] = INF
		}
	}

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		if L[i] < R[i] && L[i] < INF {
			buf.WriteByte('L')
		} else if L[i] > R[i] && R[i] < INF {
			buf.WriteByte('R')
		} else {
			buf.WriteByte('.')
		}
	}
	return buf.String()
}
