package conversions

import "fmt"

// 确保某个包中实现的类型一定满足该接口
// 不过请不要为满足接口就将它用于任何类型。
// 作为约定，仅当代码中不存在静态类型转换时才能这种声明，毕竟这是种罕见的情况。
var (
	_ IDog    = (*Dog)(nil)
	_ IAnimal = (*Animal)(nil)
	_ IObject = (*Object)(nil)
)

type Dog struct {
}

func (t *Dog) WangWang() {
	fmt.Println("Wang wang...")
}

type Animal Dog

func (t *Animal) Run() {
	fmt.Println("Run run...")
}

type Object Animal

func (t *Object) Work() {
	fmt.Println("Work work...")
}

func (t *Object) Name() {
	fmt.Println("this is object!")
}
