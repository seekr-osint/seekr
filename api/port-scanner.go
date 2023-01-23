package api

import (
	"fmt"
	"net"
	"sync"
)

func PortScan(ip string) map[int]bool { // FIXME
	wg := sync.WaitGroup{}
	valid := make(map[int]bool)
	for i := 1; i <= 65535; i++ { // 65535
		wg.Add(1)
		valid[i] = scan(&wg, ip, i)
		wg.Done()
	}

	wg.Wait()
	return valid
}

func scan(wg *sync.WaitGroup, address string, port int) bool {
	defer wg.Done()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
