package handlers

import (
	"fmt"
	"os/exec"
	"strconv"
)

func handlersping(ip string, cnt int, ch chan string) {
	count := strconv.Itoa(cnt)
	output, err := exec.Command("ping", "-c1", "-t", count, ip).Output()
	if err != nil {
		fmt.Println("Error while pinging", err)
	}
	//fmt.Println(string(output))
	ch <- string(output)
}
