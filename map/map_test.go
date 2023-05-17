package map_test

import (
	"fmt"
	"testing"
)

func Test_ModifyWithLoop(t *testing.T) {

}

func Test_DeleteElementAtAnotherFunc(t *testing.T) {
	m := map[int32]int32{1: 11, 2: 22}
	deleteElement(m)
	fmt.Println(m)
}

func deleteElement(m map[int32]int32) {
	delete(m, 1)
}

func Test_AppendKey2Slice(t *testing.T) {
	var list []*int32
	m := map[int32]int32{1: 11, 2: 22}
	for k, v := range m {
		list = append(list, &k)
		list = append(list, &v)
	}
	for _, i := range list {
		fmt.Println(*i)
	}
}

func Test_NilMapLen(t *testing.T) {
	m := make(map[int][]int)
	m[1] = []int{1, 2, 3}
	fmt.Println(len(m[1]))
	fmt.Println(len(m[2]))
}

func TestMapFreq(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
