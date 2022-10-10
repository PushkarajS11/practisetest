package main

import(
 
	"plugin"

)


func main() {
	p, err := plugin.Open("Stringconcat.so")
	if err != nil {
		panic(err)
	}
	// Str1, err := p.Lookup("Str1")
	// if err != nil {
	// 	panic(err)
	// }
	// Str2, err := p.Lookup("Str2")
	//  if err != nil {
	//  	panic(err)
	//  }
    
	Addstr, err := p.Lookup("Addstr")
	if err != nil {
		panic(err)
	}

	//*Str1.(*string) = "Hello"
	//*Str2.(*string) = "Go"

	Addstr.(func(string, string))("Hello", "go")

	//Addstr.(func())() 
}