package solution

import (
	"fmt"
	"testing"
)

//TestExample1 aaa
func TestExample1(t *testing.T) {
	nums := []int{6, 6, 0, 4, 8, 7, 6, 4, 7, 5}
	sr := Constructor()
	for _, num := range nums {
		sr.Addnum(num)
		is := sr.Getintervals()
		fmt.Printf("%v\n", is)
	}
}
