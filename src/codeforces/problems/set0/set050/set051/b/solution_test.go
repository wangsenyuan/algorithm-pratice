package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, expr string, expect []int) {
	res := solve(expr)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	expr := "<table><tr><td></td></tr></table>"
	expect := []int{1}
	runSample(t, expr, expect)
}

func TestSample2(t *testing.T) {
	expr := `<table>
	<tr>
	<td>
	<table><tr><td></td></tr><tr><td></
	td
	></tr><tr
	><td></td></tr><tr><td></td></tr></table>
	</td>
	</tr>
	</table>`
	expect := []int{1, 4}
	runSample(t, expr, expect)
}

func TestSample3(t *testing.T) {
	expr := `<table><tr><td>
	<table><tr><td>
	<table><tr><td>
	<table><tr><td></td><td></td>
	</tr><tr><td></td></tr></table>
	</td></tr></table>
	</td></tr></table>
	</td></tr></table>`
	expect := []int{1, 1, 1, 3}
	runSample(t, expr, expect)
}

func TestSample4(t *testing.T) {
	expr := `<table><tr><td><table><tr><td></td><td></td></tr></table></td><td><table><tr><td></td></tr></table></td></tr></table>`
	expect := []int{1, 2, 2}
	runSample(t, expr, expect)
}
