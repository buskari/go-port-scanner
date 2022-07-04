package port

import (
	"log"
	"net"
	"strconv"
	"time"
)

func ScanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return false
	}

	defer conn.Close()
	return true
}

func InitialScan(protocol, localhost string) {
	count := 0
	for port := 0; port <= 49152; port++ {
		if ScanPort("tcp", localhost, port) {
			log.Printf("Port opened: %v", port)
			count++
		}
	}
	if count == 0 {
		log.Println("All ports are closed")
	}
}
