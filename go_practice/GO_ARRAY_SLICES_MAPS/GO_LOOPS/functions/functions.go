package main

import "fmt"

func add(num ...int)(int){

	a:=0

	for _,v:=range num{
     a=a+v
	}

	return a
}

func reversest(str string) {


for i:=len(str)-1;i>=0; i--{
 fmt.Print(string (str[i]))

}


}




func main(){
s:= add(1,2,3,4)
fmt.Println(s)

reversest("Hello")


}