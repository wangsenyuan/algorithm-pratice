package p223

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	return (C-A)*(D-B) + (G-E)*(H-F) - commonArea(A, B, C, D, E, F, G, H)
}

func commonArea(A, B, C, D, E, F, G, H int) int {
	if A >= G || E >= C || B >= H || F >= D {
		return 0
	}
	X := max(A, E)
	W := min(C, G)
	Y := max(B, F)
	Z := min(D, H)

	return (W - X) * (Z - Y)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
