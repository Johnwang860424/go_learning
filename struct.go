package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	person1 := Person{"John", 20}
	fmt.Println(person1.name, person1.age)
	var person2 Person = Person{age: 30, name: "Jane"}
	fmt.Println(person2.name, person2.age)
	person2.name = "Alice"
	fmt.Println(person2.name, person2.age)
}
