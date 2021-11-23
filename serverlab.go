package main

import (
	"bufio"
	"fmt"
	"net"
)

func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
}


func main(){
	fmt.Println("Serverul a pornit....")

	/*listen, _ :=net.Listen("tcp",":45600")

	conexiune, _ :=listen.Accept()

	for{
		mesaj, _ :=bufio.NewReader(conexiune).ReadString('\n')
		fmt.Print("Mesajul primit ->",string(mesaj))
	}

	 */

	ln, err := net.Listen("tcp", ":45600")
	check(err, "Server is ready.")

	for {
		conn, err := ln.Accept()
		check(err, "Accepted connection.")

		go func() {
			buf := bufio.NewReader(conn)

			for {
				name, err := buf.ReadString('\n')

				if err != nil {
					fmt.Printf("Client disconnected.\n")
					break
				}

				conn.Write([]byte("Hello, " + name))
				fmt.Print("Mesajul primit ->",string(name))
			}
		}()
	}
}
