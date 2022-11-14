package for_test

import (
	"fmt"
	"testing"
)

// 编译过程： 范围循环 => 经典循环 => SSA中间代码 => 汇编机器码
// SSA https://oftime.net/2021/02/14/ssa/

// 经典循环
//
//		初始化循环的 Ninit
//		循环的继续条件 Left
//		循环体结束时执行的 Right
//		循环体 NBody
//			for Ninit; Left; Right {
//	 	  	NBody
//			}
func Test_GeneralFor(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Print(i) // 0123456789
	}
}

//		范围循环 编译器会在编译期间将所有 for-range 循环变成经典循环
//		关键字 range 帮助我们快速遍历数组、切片、哈希表以及 Channel 等集合类型
//	 经过优化的 for-range 循环的汇编代码有着相同的结构。无论是变量的初始化、循环体的执行还是最后的条件判断都是完全一样的
//		所有的 for-range 循环都会被 cmd/compile/internal/gc.walkrange 转换成不包含复杂结构、只包含基本表达式的语句。
func Test_ForRange(t *testing.T) {
	arr := []int{1, 2, 3}
	for i, _ := range arr {
		fmt.Print(i) // 012
	}
}

// 在遍历切片时追加元素 不会增加循环的执行次数
//
//	对于所有的 range 循环，Go 语言都会在编译期将原切片或者数组赋值给一个新变量 ha，
//	在赋值的过程中就发生了拷贝，而我们又通过 len 关键字预先获取了切片的长度，所以在循环中追加新的元素也不会改变循环执行的次数
func Test_ForAndAppend(t *testing.T) {
	arr := []int{1, 2, 3}
	for _, v := range arr {
		arr = append(arr, v)
	}
	fmt.Println(arr) // [1 2 3 1 2 3]
}

// 遍历slice保存变量地址到另一个slice/map时 可能出现的问题
//
//	遇到这种需要遍历元素的range循环时，Go 语言会额外创建一个新的v2变量存储切片中的元素，
//	循环中使用的这个变量v2会在每一次迭代被重新赋值而覆盖，赋值时也会触发拷贝。（v1是index）
//	因为在循环中获取返回变量的地址都完全相同，所以会发生神奇的指针现象。
//	因此当我们想要访问数组中元素所在的地址时，不应该直接获取 range 返回的变量地址 &v2，而应该使用 &a[index] 这种形式。
//	不管v2是什么类型，只要在循环中对其取地址，都是同一个地址
//	如果v2是值类型，对其所有操作仅是对当前变量v2的操作，对其拷贝的数组元素无任何作用
//	如果v2是指针类型，对其所有操作都是对其实际指向的数组元素的操作，可以修改元素的值，可以append等
func Test_ForAndAppendWithPointer(t *testing.T) {
	arr := []int{1, 2, 3}

	fmt.Println("use &value:")
	newArr1 := make([]*int, 0, len(arr))
	for _, v := range arr {
		// 这里v是int类型 所以遍历时会把每个元素的值拷贝给v
		// &v的地址是v的指针地址 不是数组中当前元素指针的地址
		fmt.Println(v) // 1 2 3
		newArr1 = append(newArr1, &v)
	}
	for _, v := range newArr1 {
		fmt.Println(*v) // 3 3 3
	}

	fmt.Println("use &arr[index]:")
	newArr2 := make([]*int, 0, len(arr))
	for i := range arr {
		newArr2 = append(newArr2, &arr[i])
	}
	for _, v := range newArr2 {
		fmt.Println(*v) // 1 2 3
	}

	fmt.Println("use []*int v")
	newArr3 := make([]*int, 0, len(arr))
	for _, v := range newArr2 {
		// 首先v是个指针 因此遍历数组元素是把每个元素地址赋值给v， 而不是把每个元素的值拷贝给v
		// v被赋值为value的指针地址 指针是地址值传递
		// 这里的v就相当于 v 指向了-> &newArr2[i] 它把当前数组元素的实际地址赋值给v了
		fmt.Println(v) // 0xc00000e198 0xc00000e1a0 0xc00000e1a8
		// 因为都是值传递
		// 这里如果用&v，就会把v的地址append到newArr里，v是复用的，每次循环地址都相同，导致结果重复
		// 而直接用v的话，就是对v的值value地址进行拷贝，然后append
		newArr3 = append(newArr3, v)
	}
	for _, v := range newArr3 {
		fmt.Println(*v) // 1 2 3
	}
}

type IPerson interface {
	GetName() string
}

type Person struct {
	Name string
}

func (p *Person) GetName() string {
	return p.Name
}

func Test_ForAndAppendWithInterface(t *testing.T) {
	persons := []IPerson{&Person{Name: "A"}, &Person{Name: "B"}, &Person{Name: "C"}}
	newPersons := make([]IPerson, 0, len(persons))
	for _, person := range persons {
		newPersons = append(newPersons, person)
	}

	for _, person := range newPersons {
		fmt.Println(person.GetName())
	}
}

// 遍历清空slice/map 直接使用for循环 编译器会清空slice/map中的数据
// cmd/compile/internal/gc.arrayClear 会优化 Go 语言遍历数组或者切片并删除全部元素的逻辑
// Go语言会直接使用 runtime.memclrNoHeapPointers 或者 runtime.memclrHasPointers 清除目标数组内存空间中的全部数据，并在执行完成后更新遍历数组的索引
func Test_CleanArr(t *testing.T) {
	arr := []int{1, 2, 3}
	for i := range arr {
		arr[i] = 0
	}

	m := map[string]int{"1": 1, "2": 2, "3": 3}
	for k := range m {
		delete(m, k)
	}
}

// 随机遍历
// Go 语言在运行时为哈希表的遍历引入了不确定性, 每次遍历顺序都是随机的, 程序不要依赖于哈希表的稳定遍历
func Test_ForMap(t *testing.T) {
	m := map[string]int{"1": 1, "2": 2, "3": 3}
	for k := range m {
		fmt.Println(k)
	}
}
