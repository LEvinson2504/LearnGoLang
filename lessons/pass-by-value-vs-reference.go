// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) person {
	return person{name, age}
}

func newPersonPointer(name string, age int) *person {
	return &person{name, age}
}

func main() {
	func() {
		fmt.Println("Pass by value")
		var person1 person = newPerson("levinson", 23)
		var person2 person = person1
		person2.age = 24
		fmt.Println(person1)
		fmt.Println(person2)
	}()

	func() {
		fmt.Println("Pass by reference")
		var person1 *person = newPersonPointer("levinson", 23)
		var person2 *person = person1
		person2.age = 24
		fmt.Println(person1)
		fmt.Println(&person1)
		fmt.Println(person2)
	}()
}

// Output
// Pass by value
// {levinson 23}
// {levinson 24}
// Pass by reference
// &{levinson 24}
// 0xc000100020
// &{levinson 24}

// Program exited.
