package goScan

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"net"
	"strconv"
	"sync"
	"time"
)

// Basic fucnto to scan a specified port. Use a net.Conn as output.
func scanPort(connMain net.Conn,port int, dialer net.Dialer, targetIP, mode string) {
	conn, err := dialer.Dial(mode, targetIP+":"+strconv.Itoa(port))
	var rsl int

	if err != nil {
		rsl = -1
	} else {
		conn.Close()
		rsl = port
		fmt.Fprint(connMain,rsl, " : ", "open")
	}
}

// Scan function. net.Conn is linked as output. It's used in a reverseShell
func Scan(conn net.Conn,targetIP, mode string) {

	// The number of I/O tcp sockets opened simultaneously is os limited. Try 'ulimit -a' cmd on linux. This semaphore is used as a limiter
	ctx := context.TODO()
	// Use ulimit in unix to set up a better semaphore value.
	sem := semaphore.NewWeighted(int64(250))
	start := time.Now()
	// Set up a timeout to avoid closing the port scan before getting the response
	dialer := net.Dialer{Timeout: 3 * time.Second}
	var waitGroup sync.WaitGroup

	// Scan the first 1000 port
	for i := 1; i < 1000; i++ {
		sem.Acquire(ctx, 1)
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			defer sem.Release(1)
			scanPort(conn,i, dialer, targetIP, mode)
		}(i)
	}
	waitGroup.Wait()
	duration := time.Since(start)
	fmt.Fprintln(conn,"scan in :", duration)
}