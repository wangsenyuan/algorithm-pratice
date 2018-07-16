package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	f, err := os.Open("./C-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var tc int
	fmt.Fscanf(f, "%d", &tc)

	for i := 1; i <= tc; i++ {
		var n int
		fmt.Fscanf(f, "%d", &n)
		ships := make([][]int, n)
		for j := 0; j < n; j++ {
			ships[j] = make([]int, 4)
			fmt.Fscanf(f, "%d %d %d %d", &ships[j][0], &ships[j][1], &ships[j][2], &ships[j][3])
		}
		res := solve(n, ships)
		fmt.Printf("Case #%d: %.7f\n", i, res)
	}
}

func solve(n int, ships [][]int) float64 {
	A := func(Y float64) float64 {
		var res float64 = math.MinInt32
		for i := 0; i < n; i++ {
			tmp := float64(ships[i][0]+ships[i][1]+ships[i][2]) - float64(ships[i][3])*Y
			if tmp > res {
				res = tmp
			}
		}
		return res
	}

	B := func(Y float64) float64 {
		var res = math.MaxFloat64
		for i := 0; i < n; i++ {
			tmp := float64(ships[i][0]+ships[i][1]+ships[i][2]) + float64(ships[i][3])*Y
			if tmp < res {
				res = tmp
			}
		}
		return res
	}

	C := func(Y float64) float64 {
		var res float64 = math.MinInt32
		for i := 0; i < n; i++ {
			tmp := float64(ships[i][0]+ships[i][1]-ships[i][2]) - float64(ships[i][3])*Y
			if tmp > res {
				res = tmp
			}
		}
		return res
	}

	D := func(Y float64) float64 {
		var res = math.MaxFloat64
		for i := 0; i < n; i++ {
			tmp := float64(ships[i][0]+ships[i][1]-ships[i][2]) + float64(ships[i][3])*Y
			if tmp < res {
				res = tmp
			}
		}
		return res
	}

	E := func(Y float64) float64 {
		var res float64 = math.MinInt32
		for i := 0; i < n; i++ {
			tmp := float64(ships[i][0]-ships[i][1]+ships[i][2]) - float64(ships[i][3])*Y
			if tmp > res {
				res = tmp
			}
		}
		return res
	}

	F := func(Y float64) float64 {
		var res = math.MaxFloat64
		for i := 0; i < n; i++ {
			tmp := float64(ships[i][0]-ships[i][1]+ships[i][2]) + float64(ships[i][3])*Y
			if tmp < res {
				res = tmp
			}
		}
		return res
	}

	G := func(Y float64) float64 {
		var res float64 = math.MinInt32
		for i := 0; i < n; i++ {
			tmp := float64(-ships[i][0]+ships[i][1]+ships[i][2]) - float64(ships[i][3])*Y
			if tmp > res {
				res = tmp
			}
		}
		return res
	}

	H := func(Y float64) float64 {
		var res = math.MaxFloat64
		for i := 0; i < n; i++ {
			tmp := float64(-ships[i][0]+ships[i][1]+ships[i][2]) + float64(ships[i][3])*Y
			if tmp < res {
				res = tmp
			}
		}
		return res
	}

	check := func(Y float64) bool {
		a := A(Y)
		b := B(Y)
		c := C(Y)
		d := D(Y)
		e := E(Y)
		f := F(Y)
		g := G(Y)
		h := H(Y)
		/*
					A ≤ x + y + z ≤ B
			   C ≤ x + y - z ≤ D
			   E ≤ x - y + z ≤ F
			   G ≤ -x + y + z ≤ H
		*/
		if a > b || c > d || e > f || g > h {
			return false
		}
		/*
			A - x ≤ y + z ≤ B - x
			   G + x ≤ y + z ≤ H + x
			   C - x ≤ y - z ≤ D - x
			   -F + x ≤ y - z ≤ -E + x
		*/
		ah := (a - h) / 2
		bg := (b - g) / 2
		if ah > bg {
			return false
		}
		ce := (c + e) / 2
		df := (d + f) / 2
		if ce > df {
			return false
		}
		if ce > bg || ah > df {
			return false
		}
		return true
	}

	var left float64
	var right = math.MaxFloat64

	for math.Abs(right-left) > 1e-7 {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return left
}

func max(a, b, c float64) float64 {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}
