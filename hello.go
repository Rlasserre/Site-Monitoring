package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoring = 3
const delay = 10

func main() {

	showIntro()
	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
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
	fmt.Println("")

	return commandRead

}

func startMonitoring() {
	fmt.Println("Start monitoring ...")

	/* sites := []string{"https://random-status-code.herokuapp.com/", */
	/* 	"https://globo.com", "https://uol.com.br/", "https://atarde.com.br"} */

	sites := readSitesFiles()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			siteMonitoring(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}
func siteMonitoring(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if resp.StatusCode == 200 {

		fmt.Println("Site:", site, "Is running")
	} else {
		fmt.Println("Site:", site, "Is Down, Status Code:", resp.StatusCode)
	}
}

func readSitesFiles() []string {

	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error :", err)
	}

	reader := bufio.NewReader(file)

	line, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error :", err)
	}

	fmt.Println(line)

	return sites
}
