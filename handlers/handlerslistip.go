package handlers

import (
	"fmt"
	"strconv"
	"strings"
)

func Handlerslistip(ip string, cnt int) {
	fmt.Println("The given IP address is ", ip)
	//Make a buffered channel of 255 size
	ch := make(chan string, 255)
	defer close(ch)
	for i := 1; i <= 255; i++ {
		parts := strings.Split(ip, ".")
		parts[3] = strconv.Itoa(i)
		ipaddress := parts[0] + "." + parts[1] + "." + parts[2] + "." + parts[3]
		fmt.Println("The new ip address is", ipaddress)
		go handlersping(ipaddress, cnt, ch)
	}
	for i := 1; i <= 255; i++ {
		fmt.Println(<-ch)
	}

}
