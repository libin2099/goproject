package main

import (
	"fmt"
	"unsafe"
)

func main() {
	decimalVar := 1
	liternalVar := "Go is good!"
	binaryVar := true

	// 指针定义1：明确指定指针的类型并赋值
	var pointer1 *int = &decimalVar
	// 指针定义2：根据值推导出指针的类型并赋值
	var pointer2 = &liternalVar
	// 指针定义3：简写省去var，并根据值推导出指针的类型并赋值
	pointer3 := &binaryVar
	// 定义一个指针，但并不初始化，其值默认为nil。
	var charactorVar *byte
	// 定义一个二级指针
	var charactorPointerVar **byte = &charactorVar

	fmt.Printf("pointer1 = %p , pointer2 = %p , pointer3 = %p \n", pointer1, pointer2, pointer3)
	fmt.Printf("*pointer1 = %d , *pointer2 = %s , *pointer3 = %t \n", *pointer1, *pointer2, *pointer3)
	// 指针未初始化，引用它指向的值时会报错, 因为引用了nil会出错
	// fmt.Println(*charactorVar)
	// 指针未初始化，初始值为nil
	fmt.Println(charactorVar)
	// 二级指针指向的地址，即使charactorVar指向了nil, 但是charactorVar依然有地址
	fmt.Println(charactorPointerVar)
	fmt.Println("---------------------------------------")
	decimalVar1 := 100
	p := &decimalVar1
	pp := &p
	fmt.Printf("decimalVar1 = %d , p = %p , pp = %p , *p = %d *pp = %p , **pp = %d \n", decimalVar1, p, pp, *p, *pp, **pp)
	// 用指针来修改值，会影响原有变量的值
	fmt.Println("---------------------------------------")
	*p = 20
	fmt.Printf(" *p = %d , decimalVar1 = %d , **pp = %d , *pp = %p , p = %p \n", *p, decimalVar1, **pp, *pp, p)
	**pp = 30
	fmt.Printf(" **pp = %d , decimalVar1 = %d , *p = %d , *pp = %p , p = %p \n", *p, decimalVar1, *p, *pp, p)
	fmt.Println("---------------------------------------")
	// 指针是不能直接运算的，如果要运算，先转换成unsafe.Pointer, 再转为uintptr进行计算，如下
	//p = p + 1  // 是错误的
	fmt.Printf("p = %p \n", p)
	up := unsafe.Pointer(p)
	uptr1 := uintptr(up)

	// 转为uintptr后可以进行计算了,这里(*uint8)(unsafe.Pointer(uptr1))表示将uptr1计算后的值转为一个地址(用(*uint8)强转为地址)，
	// 很明显，(*uint8)(unsafe.Pointer(uptr1))计算后的地址值与p的地址值相差1, 符合预期。
	uptr1 = uptr1 + 1
	fmt.Printf("uptr1 = %d , (*uint8)(unsafe.Pointer(uptr1)) = %p , p = %p \n", uptr1, (*uint8)(unsafe.Pointer(uptr1)), p) // (*uint8)转为指针类型，不是取值表达式，用括号的原因是*的优先级低，需要用括号使*与uint8结合

}
