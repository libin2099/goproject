package main

import (
	"fmt"
)

func main() {
	literal2SliceAscii()
	fmt.Println("---------------------")
	literal2SliceUnicode()
}

/*
*

	ascii的字符串和rune切片的转换，一个字符对应一个rune，不需特殊处理(长度就是字符串的长度)
*/
func literal2SliceAscii() {
	s := "This is the first go program!"
	lens := len(s)
	slice1 := s[0:len(s)]
	len1 := len(slice1)
	var slice2 []rune = []rune(s[0:len(s)])
	len2 := len(slice2)
	fmt.Println(slice1)
	fmt.Println(string(slice2))
	fmt.Println(lens == len1)
	fmt.Println(lens == len2)
	fmt.Println(len1 == len2)
}

/*
*

	对于包括了unicode的字符的字符串，一个unicode字符是三个字节，所以下面的s的
	长度是26，转为字节(或字符，2^8取值范围)切片的长度也就是26了(一个unicode转为三个字节)，但是转为rune切片后的长度为10,因为rune是
	四个字节，其值的范转是能容纳unicode字符的，所以一个unicode字符就转为一个rune,
	那么有多少个unicode就有多少个rune，所以转为rune切片后的长度为10。
*/
func literal2SliceUnicode() {
	s := "这是第一个go程序！"
	lens := len(s)
	slice1 := s[0:len(s)]
	len1 := len(slice1)
	var slice2 []rune = []rune(s[0:len(s)])
	len2 := len(slice2)
	fmt.Println(slice1)
	fmt.Println(string(slice2))
	fmt.Println(lens, len1, lens == len1)
	fmt.Println(lens, len2, lens == len2)
	fmt.Println(len1, len2, len1 == len2)
}
