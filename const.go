package main

import (
	"fmt"
)

/*
		go 的局部常量和全局常量：

	    如果go的全局常量首字母是大写，表示是公开的，即所有包的任何地方都可以访问它。
        如果是小写的全局常量，只能是包内可以访问，其它的包是不可以访问的。
        常量和变量的定义类似，除了var换成const，常量在声明时必须赋值。
*/
// 1.定义常量第一种，指定类型
const decimalConst int = 1

// 2.定义常量第二种，指定类型也赋值。
const liternalConst string = ""

// 3.定义常量第三种，不指定类型赋值，用值推导出类型
const binaryConst = true

// 4 一次定义多个常量
const (
	const1                 int = 2
	const2                     = "Go Go Go!"
	const3                     = false
	const4, const5, const6 int = 4, 5, 6
	const7, const8, const9     = false, "Golang", 4.7
)

// 给string定义一个别名，用于星期值的枚举
type WEEK string

/*
Golang没有内置的枚举类型，用const来模拟
*/
const (
	MONDAY    WEEK = "Monday"
	TUESDAY   WEEK = "Tuesday"
	WEDNESDAY WEEK = "Wednesday"
	THURSDAY  WEEK = "Thursday"
	FRIDAY    WEEK = "Friday"
	SATURDAY  WEEK = "Saturday"
	SUNDAY    WEEK = "Sunday"
)

type MONTH int

/*
*
用iota来初始化枚举的值，如果iota在第一个常量位置，那么它就是0，如果iota在第二个常量位置，那么它就是1，以此类推
*/
const (
	JAN MONTH = 1 + iota
	FEB
	MAR
	APR
	MAY
	JUN
	JUL
	AUG
	SEP
	OCT
	NOV
	DEC
)

func (w *WEEK) getWeekName() string {
	switch *w {
	case MONDAY:
		return "Today is " + string(MONDAY)
	case TUESDAY:

		return "Today is " + string(TUESDAY)
	case WEDNESDAY:
		return "Today is " + string(WEDNESDAY)
	case THURSDAY:
		return "Today is " + string(THURSDAY)
	case FRIDAY:
		return "Today is " + string(FRIDAY)
	case SATURDAY:
		return "Today is " + string(SATURDAY)
	case SUNDAY:
		return "Today is " + string(SUNDAY)
	default:
		return "Today is unknown day"
	}
}

func (m *MONTH) whichMonth() string {
	switch *m {
	case JAN:
		return "一月"
	case FEB:
		return "二月"
	case MAR:
		return "三月"
	case APR:
		return "四月"
	case MAY:
		return "五月"
	case JUN:
		return "六月"
	case JUL:
		return "七月"
	case AUG:
		return "八月"
	case SEP:
		return "九月"
	case OCT:
		return "十月"
	case NOV:
		return "十一月"
	case DEC:
		return "十二月"
	default:
		return "不知道几月"
	}
}

func main() {
	fmt.Println(decimalConst)
	fmt.Println(liternalConst)
	fmt.Println(binaryConst)
	fmt.Println(const1)
	fmt.Println(const2)
	fmt.Println(const3)
	fmt.Println(const4)
	fmt.Println(const5)
	fmt.Println(const6)
	fmt.Println(const7)
	fmt.Println(const8)
	fmt.Println(const9)
	fmt.Println("--------------------枚举-------------------")
	var today WEEK = SUNDAY // 2025.5.4
	fmt.Println((&today).getWeekName())
	// 上面的可以简写成如下
	fmt.Println(today.getWeekName())
	fmt.Println("--------------------iota-------------------")
	var curMonth MONTH = MAY
	fmt.Println(curMonth.whichMonth())

	// 单独使用iota，它总是0，也印证了第一行的iota的值是0。
	const iotaVar int = iota
	fmt.Println("iota的值是：", iotaVar)
}
