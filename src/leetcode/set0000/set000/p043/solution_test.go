package p043

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, num1, num2, expect string) {
	res := multiply(num1, num2)
	if res != expect {
		t.Errorf("Sample %s %s, expect %s, but got %s", num1, num2, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "23", "56", "1288")
}

func TestSample2(t *testing.T) {
	runSample(t, "123", "456", "56088")
}

func TestSample3(t *testing.T) {
	runSample(t, "999", "999", "998001")
}

func TestSample4(t *testing.T) {
	runSample(t, "2498002655", "5186493004848", "12955873296249231871440")
}

func TestSample5(t *testing.T) {
	runSample(t, "2655", "93004848", "246927871440")
}

func TestSample6(t *testing.T) {
	runSample(t, "655", "004848", "3175440")
}

func TestSample7(t *testing.T) {
	runSample(t, "657", "4941", "3246237")
}

func TestMulDirectly5(t *testing.T) {
	num1 := []int{6, 5, 7}
	num2 := []int{4, 9, 4, 1}
	expect := []int{3, 2, 4, 6, 2, 3, 7}
	res := mulDirectly(num1, num2)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", num1, num2, expect, res)
	}
}

func TestMulDirectly(t *testing.T) {
	num1 := []int{9, 9, 9}
	num2 := []int{9, 9, 9}
	expect := []int{9, 9, 8, 0, 0, 1}
	res := mulDirectly(num1, num2)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", num1, num2, expect, res)
	}
}

func TestMulDirectly2(t *testing.T) {
	num1 := []int{2, 4, 9, 8, 0, 0, 2, 6, 5, 5}
	num2 := []int{5, 1, 8, 6, 4, 9, 3, 0, 0, 4, 8, 4, 8}
	expect := []int{1, 2, 9, 5, 5, 8, 7, 3, 2, 9, 6, 2, 4, 9, 2, 3, 1, 8, 7, 1, 4, 4, 0}
	res := mulDirectly(num1, num2)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", num1, num2, expect, res)
	}
}

func TestMulDirectly3(t *testing.T) {
	num1 := []int{2, 6, 5, 5}
	num2 := []int{9, 3, 0, 0, 4, 8, 4, 8}
	expect := []int{2, 4, 6, 9, 2, 7, 8, 7, 1, 4, 4, 0}
	res := mulDirectly(num1, num2)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", num1, num2, expect, res)
	}
}

func TestMulDirectly4(t *testing.T) {
	num1 := []int{6, 5, 5}
	num2 := []int{0, 0, 4, 8, 4, 8}
	expect := []int{3, 1, 7, 5, 4, 4, 0}
	res := mulDirectly(num1, num2)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", num1, num2, expect, res)
	}
}
func TestAdd1(t *testing.T) {
	a := []int{2, 3}
	b := []int{1}
	c := add(a, b)
	if !reflect.DeepEqual(c, []int{2, 4}) {
		t.Errorf("Sample %v %v, expect [2, 4] but got %v", a, b, c)
	}
}

func TestSub1(t *testing.T) {
	a := []int{2, 3}
	b := []int{4}
	c := sub(a, b)
	if !reflect.DeepEqual(c, []int{1, 9}) {
		t.Errorf("Sample %v %v, got wrong answer %v", a, b, c)
	}
}
