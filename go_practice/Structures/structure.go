package main

import "fmt"

type Person struct {
	name       string
	address    string
	contact_no uint
}

func main() {

	//var p Person
	//fmt.Println(p)

	p := &Person{"Pushkaraj", "Pune", 9405845880}

	fmt.Println("Person1:", p)
	fmt.Println("Person1:", &p)
	fmt.Println("Person1:", *p)

}
