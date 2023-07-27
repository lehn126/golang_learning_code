package p2 // 包名和文件夹名不需要相同，但是建议用相同的名称

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/cast"
)

func P2_test1(name string) string {
	sb := strings.Builder{}
	sb.WriteString("hello [")
	sb.WriteString(name)
	sb.WriteString("] from p2 test1")
	return sb.String()
}

func changeArray(array [3]int, index int, value int) {
	array[index] = value
}

func changeArrayByPointer(arrayPrt *[3]int, index int, value int) {
	arrayPrt[index] = value
}

func P2_ArrayTest() {
	var array1 [3]int = [3]int{1, 2, 3}
	var array2 = [3]int{4, 5, 6}
	array3 := [3]int{10, 11, 12}

	var array4 = []int{7, 8, 9}
	array5 := []int{10, 11, 12}
	array4 = append(array4, 10)
	array5 = append(array5, 13)

	fmt.Printf("array1 = %v, type is %T, size is %v\n", array1, array1, len(array1))
	fmt.Printf("array2 = %v, type is %T\n", array2, array2)
	fmt.Printf("array3 = %v, type is %T\n", array3, array3)
	fmt.Printf("array4 = %v, type is %v\n", array4, reflect.TypeOf(array4))
	fmt.Printf("array5 = %v, type is %v\n", array5, reflect.TypeOf(array5))
	fmt.Println("element 2 in array1 is", array1[1])

	for i := 1; i < len(array2); i++ {
		fmt.Println("array2[" + strconv.Itoa(i) + "]=" + strconv.Itoa(array2[i]))
	}

	for i, v := range array2 {
		fmt.Println("array2[" + cast.ToString(i) + "]=" + cast.ToString(v))
	}

	num := 0
	for num < len(array3) {
		fmt.Println("array3[" + strconv.Itoa(num) + "]=" + strconv.Itoa(array3[num]))
		num++
	}

	ss := []string{"a", "b", "c", "d"}
	testArray(ss)

	changeArray(array3, 0, 20)
	fmt.Printf("changed array3 is %v\n", array3) // 原数组没有被修改
	changeArrayByPointer(&array3, 0, 20)
	fmt.Printf("changed array3 by pointer is %v\n", array3) // 原数组被修改了
}

func testArray(ss []string) {
	for i, v := range ss {
		fmt.Printf("index[%v] is [%v]\n", i, v)
	}
}

type MyInterface1 interface {
	SayHello() string
}

type MyInterface2 interface {
	SayBye() string
}

type StructA struct {
	Name string
}

func (sap StructA) SayHello() string {
	return sap.Name + " say hello !!!"
}

func (sa StructA) SayBye() string {
	return sa.Name + " say goodbye !!!"
}

// func ShowType(it any) {
func ShowType(it interface{}) {
	switch it.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("boolean")
	default:
		fmt.Println("unknow")
	}
}

func P2_TestInterface() {
	sa := StructA{Name: "SA"}
	defer fmt.Printf("p2 test interface say hello: %v\n", sa.SayHello())
	defer fmt.Printf("p2 test interface say bye: %v\n", sa.SayBye())
	fmt.Printf("p2 test interface struct type is: %T\n", sa)
	fmt.Printf("p2 test interface object type is: %T\n", "hello")

	prt := new(StructA)
	prt.Name = "new() test"
	fmt.Println(prt.SayHello())
}

func getClosure() func(string, string) (string, string) {
	return func(var1, var2 string) (string, string) {
		return var2, var1
	}
}

func callClosure(fun func(string, string) (string, string), var1 string, var2 string) (string, string) {
	return fun(var1, var2)
}

func P2_TestClosure() {
	swap := getClosure()
	rt1, rt2 := callClosure(swap, "hello", "closure")
	fmt.Printf("p2 test closure return: %v %v\n", rt1, rt2)
}

func changeSlice(slice []int, index int, value int) {
	slice[index] = value
}

func P2_TestSlice() {
	var slice1 = []int{7, 8, 9}
	slice2 := []int{10, 11, 12}

	slice1 = append(slice1, 10)
	slice2 = append(slice2, 13)

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("index %v in slice1 is %v\n", i, slice1[i])
	}
	for i, v := range slice2 {
		fmt.Printf("index %v in slice2 is %v\n", i, v)
	}

	changeSlice(slice2, 0, 20)
	fmt.Printf("changed slice2 is %v\n", slice2)

	slice3 := make([]string, 3) //初始容量为3，长度为3
	fmt.Printf("length of slice3 is %v\n", len(slice3))
	fmt.Printf("capacity of slice3 is %v\n", cap(slice3))
	slice3 = append(slice3, "a", "b", "c")
	fmt.Printf("length of slice3 is %v\n", len(slice3))
	fmt.Printf("capacity of slice3 is %v\n", cap(slice3))
	slice3 = append(slice3, "d")
	fmt.Printf("length of slice3 is %v\n", len(slice3))
	fmt.Printf("capacity of slice3 is %v\n", cap(slice3))

	slice4 := []string{}
	fmt.Printf("length of slice4 is %v\n", len(slice4))
	fmt.Printf("capacity of slice4 is %v\n", cap(slice4))
	slice4 = append(slice4, "e")
	fmt.Printf("length of slice4 is %v\n", len(slice4))
	fmt.Printf("capacity of slice4 is %v\n", cap(slice4))
	slice4 = append(slice4, "f")
	fmt.Printf("length of slice4 is %v\n", len(slice4))
	fmt.Printf("capacity of slice4 is %v\n", cap(slice4))
}

func P2_TestMap() {
	map1 := map[string]string{"a": "1", "b": "2", "c": "3"}
	for k, v := range map1 {
		fmt.Printf("map1 %v=%v\n", k, v)
	}
	fmt.Printf("map1 a=%v\n", map1["a"])

	map2 := make(map[string]string) // map可以省略初始容量参数，默认为空的map
	map2["c1"] = "cv1"
	map2["c2"] = "cv2"
	map2["c3"] = "cv3"
	fmt.Printf("map2 c1=%v\n", map2["c1"])

	map3 := map[string]string{}
	map3["a1"] = "v1"
	map3["a2"] = "v2"
	map3["a3"] = "v3"
	fmt.Printf("map3 a2=%v\n", map3["a2"])
	delete(map3, "a2")
	fmt.Printf("map3 a2=%v\n", map3["a2"])
}

func createError(flag int) (string, error) {
	switch flag {
	case 0:
		return "success", nil
	case 1:
		return "fail", errors.New("find a failed flag")
	default:
		// 可以使用fmt.Errorf创建简单的格式化error
		return "unknow", fmt.Errorf("this is an error for unknow flag: %v", flag)
	}
}

func P2_TestError() {
	s, e := createError(0)
	fmt.Printf("return [%v] and error is: %v\n", s, e)

	s, e = createError(1)
	fmt.Printf("return [%v] and error is: %v\n", s, e)

	s, e = createError(2)
	fmt.Printf("return [%v] and error is: %v\n", s, e)
}

func P2_TestPanicAndRecover() {
	fmt.Println("main begin")
	// 必须要先声明defer，否则不能捕获到panic异常
	defer func() {
		fmt.Println("defer begin")
		if err := recover(); err != nil {
			// 这里的err其实就是panic传入的内容
			fmt.Println(err)
		}
		fmt.Println("defer end")
	}()
	testPanic()
	// test中出现错误，这里开始下面代码不会再执行
	fmt.Println("main end")
}

func testPanic() {
	fmt.Println("test begin")
	panic("error")
	//这里开始下面代码不会再执行
	fmt.Println("test end")
}

func reflectsetvalue1(x interface{}) {
	value := reflect.ValueOf(x)
	if value.Kind() == reflect.String {
		value.SetString("欢迎来到W3Cschool")
	}
}

func P2_TestIota() {
	const (
		A = 100
		B = 200
		C = 300
	)
	fmt.Println(A, B, C)

	const (
		D = 100
		E
		F
	)
	fmt.Println(D, E, F)

	const (
		Const_A = iota // = iota = 0; iota++
		Const_B        // = iota; iota++
		Const_C        // = iota; iota++
		Const_D = 100  // = 100;  iota++
		Const_E        // = 100;  iota++
		Const_F = iota // = iota; iota++
		Const_G        // = iota; iota++
	)

	fmt.Println(Const_A, Const_B, Const_C, Const_D, Const_E, Const_F, Const_G)

	const (
		i = 1 << iota // iota = 0; = 1 << iota; iota++
		j = 3 << iota // = 3 << iota; = 3 << 1; iota++
		k             // = 3 << iota; = 3 << 2; iota++
		l             // = 3 << iota; = 3 << 3; iota++
	)
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}

func P2_TestReflect() {
	changeValueByPoint := func(v any) {
		value := reflect.ValueOf(v)
		if value.Kind() == reflect.Pointer {
			// 反射中使用Elem()方法获取指针所指向的值
			value.Elem().SetString("changed by point")
		}
	}
	sourceStr := "source value"
	ptr := &sourceStr
	changeValueByPoint(ptr)
	fmt.Println(sourceStr)
}
