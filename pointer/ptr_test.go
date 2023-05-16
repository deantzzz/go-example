package pointer_test

import (
	"fmt"
	"testing"
)

type Ptr struct {
	name string
}

func (p *Ptr) SetName(name string) {
	p.name = name
}

func (p Ptr) GetName() string {
	return p.name
}

func TestReceiver(t *testing.T) {
	var p *Ptr
	p.SetName("foo")
	got := p.GetName()
	fmt.Println(got) // panic: runtime error: invalid memory address or nil pointer dereference
	// 指针不能调用非指针方法
}
