package p3484

import (
	"strconv"
	"strings"
)

type Spreadsheet struct {
	sheet [][26]int
}

func Constructor(rows int) Spreadsheet {
	sheet := make([][26]int, rows+1)
	return Spreadsheet{sheet}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	r, c := get(cell)
	this.sheet[r][c] = value
}

func get(cell string) (r int, c int) {
	c = int(cell[0] - 'A')
	r, _ = strconv.Atoi(cell[1:])
	return
}

func (this *Spreadsheet) ResetCell(cell string) {
	r, c := get(cell)
	this.sheet[r][c] = 0
}

func (this *Spreadsheet) GetValue(formula string) int {
	i := strings.IndexByte(formula, '+')
	a := formula[1:i]
	b := formula[i+1:]
	var res int
	if a[0] >= 'A' && a[0] <= 'Z' {
		r, c := get(a)
		res = this.sheet[r][c]
	} else {
		res, _ = strconv.Atoi(a)
	}

	if b[0] >= 'A' && b[0] <= 'Z' {
		r, c := get(b)
		res += this.sheet[r][c]
	} else {
		tmp, _ := strconv.Atoi(b)
		res += tmp
	}

	return res
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */
