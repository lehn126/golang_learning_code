package main

import (
	"fmt"
	nettest "work1/net"
	p1 "work1/package1" // <module path>/<package path>, 使用包的别名
	p2 "work1/package2"
)

func TestConst() {
	const CS1 string = "const string1 in work1"
	const CS2 = "const string2 in work1"
}

// 结构体使用的方法接口
type NoteInterface interface {
	WriteOne(data string) string
	ReadOne() string
}

// 结构体及成员变量
type Note struct {
	Name  string
	Lines []string
}

// 结构体外部方法实现
func (note *Note) WriteOne(data string) string { // 使用 *Note 来访问结构体实例内的成员变量
	note.Name = "my note"
	note.Lines = append(note.Lines, data)
	return note.ReadOne()
}

// 结构体外部方法实现
func (note *Note) ReadOne() string { // 使用 *Note 来访问结构体实例内的成员变量
	if len(note.Lines) > 0 {
		return note.Lines[len(note.Lines)-1]
	} else {
		return ""
	}
}

// 测试结构体，接口和外部方法实现
func TestStructAndMethod() {
	note := Note{Name: "abc"}
	note.WriteOne("111")
	fmt.Printf("note.Name: %v\n", note.Name)
}

func TestBasic() {
	fmt.Println("hello golang world !!!")
	fmt.Println("result from p1: " + p1.P1_test1("golang"))
	fmt.Println("result from p2: " + p2.P2_test1("golang2"))
	p1.P1_TestStruct(p1.P1_Struct1Instance)
	//p2.P2_TestGoroutineMutex()
	p2.P2_TestGoroutineRWMutex()
	p2.P2_TestSelect()
	p1.P1_test2()
	p1.P1_TestRegex()
}

// 结构体中嵌入机制（组合）
type AType struct {
	NameA string
	BType // 只声明类型BType相当于为类型AType引入了BType内的成员变量与方法
	CType // 只声明类型CType相当于为类型AType引入了CType内的成员变量与方法
}
type BType struct {
	NameB string
}

func (b *BType) HelloB() {
	fmt.Println("helloB")
}

type CType struct {
}

func (c *CType) Hellocc() {
	fmt.Println("hellocc")
}

// 测试结构体嵌入机制的使用
func TestStructExtends() {
	a := AType{
		BType: BType{
			NameB: "b",
		},
	}
	a.HelloB()                           // 直接调用了 BType 中的方法
	fmt.Printf("a.NameB: %v\n", a.NameB) // 直接访问 BType 中的成员变量
	a.Hellocc()
	a.NameB = "a" // 直接修改 BType 中的成员变量
	fmt.Printf("a.BType.NameB: %v\n", a.BType.NameB)
}

func TestNet() {
	nettest.StartServices()
}

func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		fmt.Println("DeferClosureLoopV2, i:", &i)
		defer func() {
			println("DeferClosureLoopV1: ", i, &i)
		}()
	}
}

func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		fmt.Println("DeferClosureLoopV2, i:", &i)
		defer func(val int) {
			println("DeferClosureLoopV2, ", val, &val)
		}(i)
	}
}

func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		fmt.Println("DeferClosureLoopV3, i:", &i, ", j:", &j)
		defer func() {
			println("DeferClosureLoopV3, ", j, &j)
		}()
	}
}

func DeferClosureLoopV4() {
	array := []func(){}
	for i := 0; i < 10; i++ {
		j := i
		fmt.Println("DeferClosureLoopV4, i:", &i, ", j:", &j)
		array = append(array, func() {
			println("DeferClosureLoopV4, ", j, &j)
		})
	}
	for _, fun := range array {
		fun()
	}
}

func testFun(j int) {
	println("DeferClosureLoopV5, ", j, &j)
}

func DeferClosureLoopV5() {
	for i := 0; i < 10; i++ {
		j := i
		fmt.Println("DeferClosureLoopV5, i:", &i, ", j:", &j)
		defer testFun(j)
	}
}

func TestFunAndClosure() {
	DeferClosureLoopV1()
	DeferClosureLoopV2()
	DeferClosureLoopV3()
	//DeferClosureLoopV4()
	//DeferClosureLoopV5()

	a1 := DeferReturn()
	fmt.Println("a1: ", a1)

	a2 := DeferReturnV1()
	fmt.Println("a2: ", a2)

	a3 := DeferReturnV2()
	fmt.Println("a3: ", *a3)

	a4 := DeferReturnV3()
	fmt.Println("a4: ", a4)
}

func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()

	return a
}

func DeferReturnV1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()

	return a
}

func DeferReturnV2() *MyStruct {
	a := &MyStruct{
		name: "Jerry",
	}

	defer func() {
		a.name = "Tom"
	}()

	return a
}

type MyStruct struct {
	name string
}

func DeferReturnV3() (ret int) {
	defer func() {
		ret++
	}()

	return 0
}

func printObj(obj any) {
	println(obj)
}

func main() {
	//TestBasic()
	//TestNet()
	//TestFunAndClosure()
	TestStructAndMethod()
	TestStructExtends()
}
