package bit_operation_test

import (
	"fmt"
	"testing"
)

//const (
//	SubToSub = 1 * *1
//)

func TestBitOp(t *testing.T) {
	flag := 0000
	fmt.Printf("%b : %d \n", flag, flag)
	if true {
		flag = flag | 1<<3
	}
	fmt.Printf("%b : %d \n", flag, flag)
	if true {
		flag = flag | 1<<2
	}
	fmt.Printf("%b : %d \n", flag, flag)
	if true {
		flag = flag | 1<<1
	}
	fmt.Printf("%b : %d \n", flag, flag)
	if true {
		flag = flag | 1<<0
	}
	fmt.Printf("%b : %d \n", flag, flag)

	fmt.Println()
}
