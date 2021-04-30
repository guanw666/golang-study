package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("randomInt", rand.Intn(10))
	fmt.Printf("a %g b\n", math.Sqrt(9))
	// 大写开头的为导出的
	fmt.Println(math.Pi)
	fmt.Println(add(5, 2))
	fmt.Println(swap("x", "y"))
	fmt.Println(split(10))
	fmt.Println(a, b, c)
	fmt.Println(a1, b1, c1)
	assignment()
	basicTypes()
	typeConvert()
	typeDerive()
	constTest()
	numConstantTest()
	forTest()
	ifTest()
	switchTest()
	deferTest()
	deferTest2()
	pointerTest()
	structTest()
}

// 函数结构
func add(x, y int) int {
	return x + y
}

// 函数多值返回
func swap(a, b string) (string, string) {
	return b, a
}

// 没有参数的return返回已命名的返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 声明变量列表
var a, b, c bool

// 声明同时赋值
var a1, b1, c1 = true, false, "c1"

// :=在函数体内可代替var声明
func assignment() {
	var a = 1
	b := 2
	c, d, e := true, false, "e"
	fmt.Println(a, b, c, d, e)
}

// 基本类型
// bool
// string
// int int8 int16 int32(rune) int64
// uint uint(byte) uint16 uint32 uint64 uintptr
// float32 float64
// complex64 complex128
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(15 + 12i)
)

func basicTypes() {
	fmt.Printf("Type:%T Value:%v\n", ToBe, ToBe)
	fmt.Printf("Type:%T Value:%v\n", MaxInt, MaxInt)
	fmt.Printf("Type:%T Value:%v\n", z, z)
}

// GO需要显式类型转换
func typeConvert() {
	a := 1
	b := 2.5
	fmt.Printf("Type:%T Value:%v\n", float64(a), float64(a))
	fmt.Printf("Type:%T Value:%v\n", uint(b), uint(b))
}

// 类型推导
func typeDerive() {
	var i int
	j := i
	fmt.Printf("Type:%T\n", j)
	// 当右边包含未指明类型的数值常量时，新变量的类型就可能是 int, float64 或 complex128 了，这取决于常量的精度
	b := 0.5i + 0.123
	fmt.Printf("Type:%T\n", b)
}

// const不能使用:=
const consta = 1
const constb = false
const constc = "c"

func constTest() {
	fmt.Println(consta, constb, constc)
}

// 数值常量是高精度的值
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}
func numConstantTest() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func forTest() {
	sum := 0
	for i := 1; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum", sum)
	sum1 := 1
	for ; sum1 < 100; {
		sum1 += sum1
	}
	fmt.Println("sum1", sum1)
	sum2 := 1
	// 相当于while
	for sum2 < 100 {
		sum2 += sum2
	}
	fmt.Println("sum2", sum2)
	// 无限循环
	//for {
	//	fmt.Println("无限")
	//}
}

func ifTest() {
	b := 2
	i := 2 * b
	if i < 5 {
		fmt.Println("i<5", i < 5)
	}
	// if的简短语句，作用域if体内
	if j := math.Pow(5, 2); j >= 25 {
		fmt.Println("j>=25", true)
	} else {
		fmt.Println("else")
	}
}
func switchTest() {
	// Go 自动提供了在这些语言中每个 case 后面所需的 break 语句
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("linux.")
	default:
		fmt.Printf("%s. \n", os)
	}

	// switch 的求值顺序 从上到下顺次执行，直到匹配成功时停止。
	today := time.Now().Weekday()
	fmt.Printf("Today is %s, when is Saturday?\n", today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Two far away.")
	}

	// 没有条件的 switch 同 switch true 一样。
	// 这种形式能将一长串 if-then-else 写得更加清晰。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferTest() {
	// defer 语句会将函数推迟到外层函数返回之后执行。
	defer fmt.Println("world")
	fmt.Println("hello")
}

func deferTest2() {
	// 推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	fmt.Println("Counting")
	fmt.Println("Done")
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}

func pointerTest() {
	// int类型指针
	var p *int
	i := 2
	// & 操作符会生成一个指向其操作数的指针。
	p = &i
	// 打印内存地址
	fmt.Println(p)
	// 打印值（通过指针读取i）
	fmt.Println(*p)
	// 通过指针设置i
	*p = 3
	fmt.Println(i)
}

// 结构体（struct）就是一组字段
type Coordinate struct {
	X int
	Y int
}

func structTest() {
	v := Coordinate{1, 2}
	// 结构体字段使用点号访问
	fmt.Println(v.X, v.Y)
	// 如果我们有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X。
	//不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以。
	p := &v
	(*p).X = 3
	fmt.Println((*p).X)
	p.X = 4
	fmt.Println(p.X)
}
