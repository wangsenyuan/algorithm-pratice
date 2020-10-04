package p1600

import (
	"reflect"
	"testing"
)

func TestSample1(t *testing.T) {
	king := Constructor("king")
	king.Birth("king", "andy")
	king.Birth("king", "bob")
	king.Birth("king", "catherine")
	king.Birth("andy", "matthew")
	king.Birth("bob", "alex")
	king.Birth("bob", "asha")

	res := king.GetInheritanceOrder()

	expect := []string{"king", "andy", "matthew", "bob", "alex", "asha", "catherine"}

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample fail, expect %v, but got %v", expect, res)
	}

	king.Death("bob")

	res = king.GetInheritanceOrder()

	expect = []string{"king", "andy", "matthew", "alex", "asha", "catherine"}

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample fail after bob die, expect %v, but got %v", expect, res)
	}
}
