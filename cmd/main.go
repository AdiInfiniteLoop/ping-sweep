package main

import (
	"fmt"
	Handlers "github.com/AdiInfiniteLoop/PingSweeper/handlers"
	"log"
	"net"
)

func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalln("An error occurred while getting local IP addresses", err)
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() { //get the iip with a mask(i.r.. no MAC address) with no loopback address and all ok
			if ip4 := ipnet.IP.To4(); ip4 != nil {
				return ip4.String()
			}
		}
	}
	return ""
}

func main() {
	fmt.Println("Welcome to Ping Sweeper")
	LocalIPString := GetLocalIp()
	var cnt int
	_, err := fmt.Scanf("%d", &cnt)
	if err != nil {
		log.Fatalln("Error while scanning", err)
	}

	// If successful, print the number of packets
	fmt.Println("You chose to send", cnt, "packets.")
	Handlers.Handlerslistip(LocalIPString, cnt)
}
