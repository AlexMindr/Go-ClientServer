package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)
func check2(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

type ClientJob struct {
	name string
	conn net.Conn
}

func generateResponses(clientJobs chan ClientJob) {
	for {
		// Wait for the next job to come off the queue.
		clientJob := <-clientJobs

		// Do something thats keeps the CPU buys for a whole second.
		for start := time.Now(); time.Now().Sub(start) < time.Second; {
		}

		// Send back the response.
		clientJob.conn.Write([]byte("Hello, " + clientJob.name))
	}
}

func main() {
	clientJobs := make(chan ClientJob)
	go generateResponses(clientJobs)

	ln, err := net.Listen("tcp", ":45600")
	check2(err, "Server is ready.")

	for {
		conn, err := ln.Accept()
		check2(err, "Accepted connection.")

		go func() {
			buf := bufio.NewReader(conn)

			for {
				name, err := buf.ReadString('\n')

				if err != nil {
					fmt.Printf("Client disconnected.\n")
					break
				}

				clientJobs <- ClientJob{name, conn}
			}
		}()
	}
}