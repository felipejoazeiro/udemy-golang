package main

import (
	"fmt"
	"reflect"
)

func main(){
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2","Posição 3","Posição 4","Posição 5"}
	fmt.Println(array2)
	//array2[5] = "Posição 6"

	array3 := [...]int{1,2,3,4,5}
	fmt.Println(array3)
	array3[7] = 10
	fmt.Println(array3)

	slice := [...]int{10, 11, 12 ,13, 14, 15 ,16, 17}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice,18)
	fmt.Printlkn(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	//Arrays interno
	fmt.Println("....")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}