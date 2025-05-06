package main

import (
	"fmt"
)

/*
定义结构时，可以给出成员变量的名称，也可以不给出名称只给类型，但是一种类型只能出现一次，如下所示
*/
type Object struct {
	id int
}
type Person struct {
	name   string
	age    int
	sex    bool
	height float64
	weight float64
	int
	string
}

type Student struct {
	person Person
	name   string
	score  float64
	int
	string
	Person
	Object
	id byte
}

func (s Student) setNameSelf(name string) {
	s.name = name
}

/*
  注意：在Golang中，结构体的方法是没有重载的概念的，所以下面的方法会报错，在Golang中的方法名称都是不同的。
*/
//func (s Student) setName(sign bool, firstname string, lastname string) {
//	s.name = firstname + " " + lastname
//}

func (s *Student) setNamePointer(name string) {
	s.name = name
}

func main() {
	person := Person{
		name: "libin", age: 45, sex: true, height: 1.68, weight: 69.5, int: 66, string: "King",
	}

	partner := Person{
		name: "daizuo", age: 43, sex: false, height: 1.58, weight: 46.2, int: 77, string: "Queen",
	}

	body := Object{id: 999}
	student := Student{
		person: person,
		name:   "student_libin",
		score:  777.7,
		int:    888,
		string: "KingdomHeart",
		Person: partner,
		Object: body,
		id:     1,
	}

	// 访问结构体中没有名称的成员变量时，直接用".类型"来访问，如下
	fmt.Println(person.int)
	fmt.Println(person.string)
	fmt.Println(person)

	// 直接.name访问的是student的成员name,如果想访问它的person成员对象的name成员变量，则需要用student.person.name
	fmt.Println(student.name)
	fmt.Println(student.person.name)
	// 如果想访问没有名称的结构体成员变量时，例如Person(partner)中的变量，则需要用student.Person.name
	fmt.Println(student.Person.name)
	// 但是如果结构体中有一个没有名称的另一个结构体类型的成员变量时，并且这个结构体类型中的成员变量与外层结构体的成员变量没有重名的，那么可以直接
	// 用"."的方式访问，如下;但是如果内外层都有同一个名称的成员变量，那么即使它们的类型不同，也会应用外层的(即调用者本身的变量)
	// 一句话，就是如果内外层一共只有一个名称的成员变量，可以直接用点的方式，如果有两个及以上的，如用直接用点的方式访问，那么访问的是最外层的(即调用者本身的变量)
	fmt.Println(student.id)
	fmt.Println(student.Object.id)
	fmt.Println("------------------------结构体方法-------------------------")
	student.setNameSelf("Self")
	// 可以看出setNameSelf并没有改变student的name的值，因为是将student复制了一份，改的是复制品的name的值。
	fmt.Println(student.name)
	student.setNamePointer("Pointer")
	// 可以看出setNamePointer改变了student的name的值，因为传递的是指针，改的是原来的结构体的name的值(即student的name的值)
	fmt.Println(student.name)

	student1 := student
	student1.name = "student1's name"
	// student1是student的复制品，修改的name的值，只是修改了这个复制品的name的值，原来的student的name的值没有改变
	fmt.Println(student.name)
	// 复制品的name的值被改变了(student1's name)
	fmt.Println(student1.name)

	student2 := &student
	student2.name = "student2's name"
	// student2是指向student的指针，修改name的值，原来的student的name的值也被修改了。
	fmt.Println(student.name)
	// 因为student2指向的是student,所以student2.name和student.name的值是一个，也就相同了。
	fmt.Println(student2.name)

}
