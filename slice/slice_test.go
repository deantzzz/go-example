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
