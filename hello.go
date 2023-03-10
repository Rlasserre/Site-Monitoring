package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			showLogs()
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
	version := 1.3
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
		logger(site, true)
	} else {
		fmt.Println("Site:", site, "Is Down, Status Code:", resp.StatusCode)
		logger(site, false)
	}
}

func readSitesFiles() []string {

	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error :", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}
	file.Close()

	return sites
}

func logger(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02 Jan 06 15:04 MST") +
		" - " + site + " - online: " +
		strconv.FormatBool(status) + "\n")

	file.Close()
}

func showLogs() {

	file, err := ioutil.ReadFile("log.txt")

	fmt.Println(string(file))

	if err != nil {
		fmt.Println(err)
	}

}
