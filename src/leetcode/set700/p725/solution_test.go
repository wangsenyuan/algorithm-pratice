package p725

import (
	. "leetcode/util"
	"testing"
)

func TestSample1(t *testing.T) {
	root := ParseAsList("[1,2,3]")

	// fmt.Printf("%v\n", root.String())

	res := splitListToParts(root, 5)

	if len(res) != 5 {
		t.Errorf("result should have 5 elements, but %d", len(res))
	} else {
		for i := 0; i < 3; i++ {
			elem := res[i]
			if elem == nil {
				t.Errorf("element at %d should not be nil", i)
			} else if elem.Next != nil {
				t.Errorf("element at %d should have a nil tail", i)
			}
		}
		for i := 3; i < 5; i++ {
			elem := res[i]
			if elem != nil {
				t.Errorf("element at %d should be nil", i)
			}
		}
	}
}

func TestSample2(t *testing.T) {
	root := ParseAsList("[1,2,3,4,5,6,7,8,9,10]")
	res := splitListToParts(root, 3)
	expect := []string{"[1,2,3,4]", "[5,6,7]", "[8,9,10]"}
	if len(res) != len(expect) {
		t.Errorf("should have 3 elements, but got %d", len(res))
	} else {
		for i := 0; i < 3; i++ {
			elem := res[i]
			if elem == nil {
				t.Errorf("element at %d should not be nil", i)
			} else if elem.String() != expect[i] {
				t.Errorf("element at %d should be %s, but is %s", i, expect[i], elem.String())
			}
		}
	}
}
