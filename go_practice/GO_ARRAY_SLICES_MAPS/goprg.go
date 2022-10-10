package main

import "fmt"

func main() {

	// var arr [4] string
	// arr[0]="B"
	// arr[1]="A"
	// arr[2]="C"
	// arr[3]="D"

	// fmt.Println(arr)

	// var arr1 = [3]int{1,2,3}

	// fmt.Println(arr1)

	// var ages=map[string]int{"Person1":99,"Person2":88}

	// fmt.Println(ages["Person2"])

	var arr = [3]int{1, 2, 3}
	max := 0

	for i := 0; i <= len(arr)-1; i++ {
		if max < arr[i] {

			max = arr[i]

		}

	}
	fmt.Println(max)

}
