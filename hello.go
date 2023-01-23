package main

import (
	"fmt"
	"os"
)

func main() {

	showIntro()
	showMenu()
	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Start monitoring ...")
	case 2:
		fmt.Println("Showing Logs ...")
	case 0:
		fmt.Println("Leaving program!")
		os.Exit(0)
	default:
		fmt.Println("Unavailable option.")
		os.Exit(-1)
	}
}

func showIntro() {

	name := "Rafael"
	version := 1.1
	fmt.Println("Hi, sr.", name, "!")
	fmt.Println("Program version:", version)
}

func showMenu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit Program")
}

func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)
	fmt.Println("You choose option", commandRead)

	return commandRead

}
