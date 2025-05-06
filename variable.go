package main

import (
	"fmt"
)

/*
		go 的局部变量和全局变量：

	    如果go的全局变量首字母是大写，表示是公开的，即所有包的任何地方都可以访问它。
        如果是小写的全局变量，只能是包内可以访问，其它的包是不可以访问的。
*/
// 1.定义变量第一种，指定类型并赋值
var decimalVar int = 1

// 2.定义变量第二种，指定类型不赋值，默认值为相应类型的初值(int 0,bool false, string "" 等)
var liternalVar string

// 3.定义变量第三种，不提定类型赋值，用值推导出类型
var binaryVar = true

// 4.简写，省略var且不指定类型，赋值并推导出类型，但这种写法必须是局部变量
//floatVar := float64(1.0)

// 5.1 一次定义多个变量, 方式一
var (
	var1             int = 2
	var2                 = "Go Go Go!"
	var3             bool
	var4, var5, var6 int = 4, 5, 6
	var7, var8, var9     = false, "Golang", 4.7
)

// 5.2 方式二
var intVar1, intVar2, intVar3 int = 1, 2, 3
var liternalVar1, liternalVar2, liternalVar3 string
var binaryVar1, binaryVar2, binaryVar3 = true, true, false
var multiVar1, multiVar2, multiVar3 = 1, "libin", true

// 一次定义多个变量的简写，省略var且不指定类型，赋值并推导出类型，但这种写法必须是局部变量
//shortVar1, shortVar2, shortVar3 := 100, "success", true

func main() {
	fmt.Printf(" decimalVar = %d\n", decimalVar)
	fmt.Printf(" liternalVar = %s\n", liternalVar)
	fmt.Printf(" binaryVar = %t\n", binaryVar)
	floatVar := 1.0
	fmt.Printf(" floatVar = %f\n", floatVar)
	fmt.Printf("------------以下为5.1---------------\n")
	fmt.Printf(" var1 = %d\n", var1)
	fmt.Printf(" var2 = %s\n", var2)
	fmt.Printf(" var3 = %t\n", var3)
	fmt.Printf("------------以下为5.2---------------\n")
	fmt.Printf("intVar1 = %d , intVar2 = %d , intVar3 = %d \n", intVar1, intVar2, intVar3)
	fmt.Printf("liternalVar1 = %s , liternalVar2 = %s , liternalVar3 = %s \n", liternalVar1, liternalVar2, liternalVar3)
	fmt.Printf("binaryVar1 = %t , binaryVar2 = %t , binaryVar3 = %t \n", binaryVar1, binaryVar2, binaryVar3)
	fmt.Printf("multiVar1 = %d , multiVar2 = %s , multiVar3 = %t \n", multiVar1, multiVar2, multiVar3)
	shortVar1, shortVar2, shortVar3 := 100, "success", true
	fmt.Printf("shortVar1 = %d , shortVar2 = %s , shortVar3 = %t \n", shortVar1, shortVar2, shortVar3)

}
