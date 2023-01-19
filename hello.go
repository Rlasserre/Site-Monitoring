package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name = "Rafael"
	var version float32 = 1.1
	var age = 38
	fmt.Println("My first go program!")
	fmt.Println("Hi, sr.", name, "!")
	fmt.Println("You are", age, "y.o.")
	fmt.Println("Program version:", version)

	fmt.Println(`The type of the variable "name" is`, reflect.TypeOf(name))

}
