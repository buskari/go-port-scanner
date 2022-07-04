package main

import (
	"log"

	"github.com/buskari/go-port-scanner/port"
)

func main() {
	log.Println("Scanning TCP ports...")
	port.InitialScan("tcp", "localhost")
	log.Println("Scanning UDP ports...")
	port.InitialScan("udp", "localhost")
}
