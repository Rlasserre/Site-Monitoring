package main

import (
	"fmt"
)

func main() {
	name := "Rafael"
	var version float32 = 1.1
	fmt.Println("Hi, sr.", name, "!")
	fmt.Println("Program version:", version)

	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit Program")
	var comand int
	fmt.Scan(&comand)

	fmt.Println("You choose option", comand)
}
