package main

import "fmt"

func main (){

  // var var1 bool=false;

//    if var1 {
//    fmt.Println("Pass")
//    } else {
//    fmt.Println("Fail")
//    }

//    var num1 int=2
//    var num2 int=5
//     num3:=num1+num2

//    fmt.Println(num3)

     var age int=72


	 switch{
   
	 case age>0 && age <=2:
		fmt.Println("Infant")
	 case age>2 && age <=5:
        fmt.Println("Toddler")
	 case age>5 && age <=12:
        fmt.Println("Child")
	 case age>12 && age <=18:
        fmt.Println("Teen")
	 default:
        fmt.Println("Adult")
	 }


	
}