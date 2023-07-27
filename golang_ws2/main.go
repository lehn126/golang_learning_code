package main

import (
	"fmt"
	wp1 "work1/package1" // 第三方包: <module path>/<package path>, 使用包的别名
	p2 "work2/package2"
	p3 "work2/package3"
)

func P2_tests() {
	p2.P2_ArrayTest()
	p2.P2_TestInterface()
	p2.P2_TestClosure()

	p2.P2_TestSlice()
	p2.P2_TestMap()
	p2.P2_TestError()
	p2.P2_TestPanicAndRecover()
	p2.P2_TestIota()
	p2.P2_TestReflect()
}

func P3_tests() {
	p3.TestEnv()
}

func main() {
	fmt.Println("result from work1 p1: " + wp1.P1_test1("golang"))
	fmt.Println("result from work2 p2: " + p2.P2_test1("golang"))

	fmt.Printf("public var1 in work1 P1_Struct1 is: %v\n", wp1.P1_Struct1Instance.Public_var1)
	// 小写开头的struct成员变量在包外部不可见
	//fmt.Printf("private var2 in work1 P1_Struct1 is: %v\n", wp1.P1_Struct1Instance.private_var2) // 编译错误 wp1.P1_Struct1Instance.private_var2 undefined

	//P2_tests()
	P3_tests()
}
