package main

import (
	"fmt"
)

func main() {
	num1, num2 := 10, 5
	fmt.Printf("-------------算数运算符------------")
	fmt.Printf("num1 + num2 = %d\n", num1+num2)
	fmt.Printf("num1 - num2 = %d\n", num1-num2)
	fmt.Printf("num1 * num2 = %d\n", num1*num2)
	fmt.Printf("num1 / num2 = %d\n", num1/num2)
	fmt.Printf("num1 %% num2 = %d\n", num1%num2)
	fmt.Printf("-------------位运算符------------\n")
	fmt.Printf("num1 & num2 = %d\n", num1&num2) // 1010 & 0101 = 0000     0
	fmt.Printf("num1 | num2 = %d\n", num1|num2) // 1010 | 0101 = 1111     15
	fmt.Printf("num1 ^ num2 = %d\n", num1^num2) // 1010 ^ 0101 = 1111     15
	fmt.Printf("^num1 = %d\n", ^num1)
	fmt.Printf("num1 << 1 = %d\n", num1<<1)
	fmt.Printf("num1 >> 1 = %d\n", num1>>1)
	fmt.Printf("-------------关系运算符------------\n")
	fmt.Printf("num1 > num2 = %t\n", num1 > num2)
	fmt.Printf("num1 < num2 = %t\n", num1 < num2)
	fmt.Printf("num1 >= num2 = %t\n", num1 >= num2)
	fmt.Printf("num1 <= num2 = %t\n", num1 <= num2)
	fmt.Printf("num1 == num2 = %t\n", num1 == num2)
	fmt.Printf("num1 != num2 = %t\n", num1 != num2)
	fmt.Printf("-------------逻辑运算符------------\n")
	binary1, binary2 := true, false
	fmt.Printf("binary1 && binary2 = %t\n", binary1 && binary2)
	fmt.Printf("binary1 || binary2 = %t\n", binary1 || binary2)
	fmt.Printf("!binary1 = %t\n", !binary1)

	// 括号可以改变优先级，&(取地址)和*(指针运算)会有较高的优先级
	decimal1 := 100
	p := &decimal1
	decimal2 := 120
	decimal2 = decimal2 * *p
	fmt.Printf("decimal2 = %d\n", decimal2)

	// float64转整型会把小数部分舍掉，不会进行四舍五入，直接舍掉只留整数部分
	float1 := 12.82
	decimal3 := int(float1)
	fmt.Printf("decimal3 = %d\n", decimal3)

	// golang 只有后自增自减，没有前自增自减; 自增自减只能单独使用，不能和其它老运算混用，打印也不行
	decimal4 := 18
	decimal4++
	//decimal5 := decimal4++ //不能混用
	// fmt.Printf("decimal4++ = %d\n", decimal4++) // 这种写法错误，不能自增与Printf混用
	fmt.Printf("decimal4++ = %d\n", decimal4)
	decimal4--
	fmt.Printf("decimal4-- = %d\n", decimal4)

	// 优先级,从高到低
	// * / % << >> & ^(按位取反) !(逻辑非)
	// + - | ^
	// == != < > <= >=
	// && ||

	// 括号：()。
	// 单目运算符：!, +, -, ^, *, &, <type>（类型转换）。
	// 乘法和除法：*, /, %。
	// 加法和减法：+, -, |(按位或), ^(位异或)。
	// 关系运算符：<, <=, >, >=。
	// 相等性运算符：==, !=。
	// 逻辑与（AND）：&&。
	// 逻辑或（OR）：||。
}
