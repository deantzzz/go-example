package slice_test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_StringJoin(t *testing.T) {
	s := []string{"a", "b", "c"}
	fmt.Println(s)
	fmt.Println(strings.Join(s, ","))
}

func Test_slip(t *testing.T) {
	list := []int{1, 2, 3}
	size := len(list)
	for _, i := range list {
		fmt.Println(i)
	}
	list = list[size:]
	fmt.Println(list)
}

func Test_slice_clean(t *testing.T) {

}
