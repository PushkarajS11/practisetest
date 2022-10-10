package main

import (
	"plugin"
)

func main() {
	p, err := plugin.Open("sum.so")
	if err != nil {
		panic(err)
	}
	Num1, err := p.Lookup("Num1")
	if err != nil {
		panic(err)
	}
	// Num2, err := p.Lookup("Num2")
	// if err != nil {
	// 	panic(err)
	// }
    
	Sum, err := p.Lookup("Sum")
	if err != nil {
		panic(err)
	}

	*Num1.(*int) = 7
	//*Num2.(*int) = 8

	//Sum.(func(int, int))(1, 2)

	Sum.(func())()
}
