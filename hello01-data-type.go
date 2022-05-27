package main

import "fmt"

/**
 * 基础数据类型，变量定义，常量定义，数据类型强制转换
 */

// 定义常量（注意：Go里面常量名不一定要大写）
const aa = "sadasdsadas"

// 定义常量，同时定义多个常量
const (
	bbb = 10
	ccc = 20
)

func main() {
	// 定义常量
	const aaa string = "aaaaa"
	variableType()
	variableObject()
	convertObject()
	pointer()
}

/**
 * 简单数据类型强制转换
 */
func convertObject() {
	var a1 int32 = 10
	// 强制转换
	var a2 float32 = float32(a1)
	fmt.Println(a1, a2)
}

/**
 * 基础数据类型
 */
func variableType() {
	var abool bool = false
	var abyte byte = 'a'
	// 这个就是各大语言的char（字符）类型（注意：Go里面这个类型是占4个字节的，但是它只存一个字符）
	var arune rune = 'b'
	var afloat float32 = 10.01
	var acomplex complex64 = 10101
	fmt.Println(abool, abyte, arune, afloat, acomplex)
}

/**
 * 变量定义
 */
func variableObject() {
	var a1 int32 = 1
	// 同时定义多个变量并赋值
	var a2, a3 int32 = 2, 3
	// 不指定类型直接定义变量
	var a6, a7 = "12545", 10
	var a4 string = "sfdasl"
	// 使用冒号定义变量，表示初始赋值，不需要写var
	a5 := 10
	fmt.Printf("a1=%d,a2=%d,a3=%d,a4=%q,a5=%d,a6=%q,a7=%d", a1, a2, a3, a4, a5, a6, a7)
}

func pointer() {
	var a int = 2
	// pa指针指向a变量内存地址
	var pa *int = &a
	// 修改pa指针内存地址的值，也就是a变量的值
	*pa = 3
	fmt.Println(a)
}
