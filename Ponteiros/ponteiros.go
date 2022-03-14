package main

import "fmt"

func main() {
	variavel1 := 10
	variavel2 := variavel1

	fmt.Println(variavel1, variavel2)

	variavel1 ++

	fmt.Println(variavel1, variavel2)

	var variavel3 int 
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)

	*ponteiro ++

	fmt.Println(variavel3, ponteiro)

	slice := []int{}

	fmt.Println(len(slice), cap(slice))
	slice = append(slice, 1)
	fmt.Println(len(slice), cap(slice))

}