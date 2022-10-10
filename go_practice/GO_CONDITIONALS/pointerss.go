package main

import "fmt"


func main(){

   fmt.Println("Enter a number")

  res:=0
  fmt.Scanln(&res)

//   fmt.Println("Entered number is",res )
   
   b:=&res
   fmt.Println(b)

   fmt.Println(*b)

   fmt.Println(&b)

    c:=&b

    fmt.Println(*c)

  


}