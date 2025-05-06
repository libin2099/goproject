package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := judgeBigger(6, 10)
	if result {
		fmt.Println("great than or equals to")
	} else {
		fmt.Println("less than")
	}

	/*
	  注意：在if当中定义的变量，只能在if语句的范围内使用(可以是if,后续的else, 后续的else if中都有效)，
	  即在if...else语句中的哪里定义了变量，就在if...else的整个语句体系的后续中有效。出了if...else的整个语句体系就无效了
	*/
	if num := selectDigital(100); num > 0 {
		fmt.Printf("%d is greater than zero.\n", num)
	} else {
		fmt.Printf("%d is less than or equals to zero.\n", num)
		num = 8
		fmt.Printf("now num is %d", num)
	}

	// num 不在if范围内了，所以报 undefined: num异常
	//num = 8
	//fmt.Println(num)
	fmt.Println("--------------switch 1----------------")
	fmt.Printf("Today is %s\n", convertToWeekDay(2))

	fmt.Println("--------------switch 2----------------")
	fmt.Printf("This month has %d days\n", countDays("Jul"))

	fmt.Println("--------------switch 3----------------")
	fmt.Println(compare(7, 100))

	fmt.Println("--------------switch 4----------------")
	//var param interface{} = 1
	//var param interface{} = false
	//var param interface{} = "libin"
	var param interface{} = 67.11
	fmt.Printf(judgeType(param), param)
}

func judgeBigger(num1 int, num2 int) (result bool) {
	if num1 >= num2 {
		result = true
		fmt.Printf("%d great than or equals to %d\n", num1, num2)
	} else {
		fmt.Printf("%d less than %d\n", num1, num2)
		result = false
	}

	return
}

func selectDigital(index int) (result int) {
	digits := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if index < 0 || index >= len(digits) {
		result = -1
	} else {
		result = digits[index]
	}
	return
}

func convertToWeekDay(day int) (result string) {
	switch day {
	case 0:
		result = "Sunday"
	case 1:
		result = "Monday"
	case 2:
		result = "Tuesday"
	case 3:
		result = "Wednesday"
	case 4:
		result = "Thursday"
	case 5:
		result = "Friday"
	case 6:
		result = "Saturday"
	default:
		result = "Unknown"
	}
	return
}

func countDays(month string) (days int) {
	switch curMonth := month; curMonth {
	case "Jan", "Mar", "May", "Jul", "Aug", "Oct", "Dec":
		days = 31
	case "Apr", "Jun", "Sep", "Nov":
		days = 30
	case "Feb":
		days = 28
	default:
		days = 30
	}
	return
}

func compare(num1 int, num2 int) (result string) {
	switch {
	case num1 > num2:
		result = strconv.Itoa(num1) + " greater than " + strconv.Itoa(num2)
	case num1 < num2:
		result = strconv.Itoa(num1) + " less than " + strconv.Itoa(num2)
	case num1 == num2:
		result = strconv.Itoa(num1) + " equals " + strconv.Itoa(num2)
	}
	return
}

func judgeType(someValue interface{}) (result string) {
	switch someValue.(type) {
	case int:
		result = "%d is integer"
	case byte:
		result = "%b is byte"
	case string:
		result = "%s is string"
	case float64:
		result = "%f is float64"
	case bool:
		result = "%t is bool"
	}
	return
}
