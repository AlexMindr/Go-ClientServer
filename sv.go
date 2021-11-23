package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type ii struct{
	val int64
}

type clients struct{
	name string
	conn net.Conn
}

func reverseNumber(num int64) int64 {
	var res int64
	for num>0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	return res
}

func sum(s []int64, c chan int64) {

	var sum int64
	for _, v := range s {
		if v<0{
		sum += -1*reverseNumber(v*-1)
		} else {
			sum+=reverseNumber(v)
		}
	}
	c <- sum // send sum to c
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func sub3(conn net.Conn,name string){
	dec := gob.NewDecoder(conn)
	var p []int64
	err := dec.Decode(&p)
	checkError(err)
	fmt.Printf("Received from %s the following data: %+v\n", name, p)
	newmessage := "Connected as: "+name
	_, err2 := fmt.Fprintf(conn, "%s \n", newmessage)
	checkError(err2)
	time.Sleep(1000*time.Millisecond)
	_, err8 := fmt.Fprintf(conn, "%s \n", "Request recieved")
	checkError(err8)

	//time.Sleep(5000*time.Millisecond)
	_, err3 := fmt.Fprintf(conn,"%s \n","Server is processing...")
	checkError(err3)

	time.Sleep(5000*time.Millisecond)

	c := make(chan int64)
	go sum(p, c)
	suma:=<-c


	_, err4 := fmt.Fprintf(conn,"Server sends %d as reply \n",suma)
	checkError(err4)

	messageres, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("%s\n", messageres)

	err6 := conn.Close()
	checkError(err6)
}

func handleConnection(client clients) {
	conn:=client.conn

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("Client %s connected with id%s\n", message,client.name)

	name:=message
	msg, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("%v",msg)
	subiect:=strings.TrimSuffix(msg, "\n")
	switch subiect{
	case "3":
		sub3(conn,name)
	case "5":\b[01]+\b
	default:
		fmt.Println("Nu exista subiectul!")
	}
	fmt.Printf("Closing connection\n")
}


func main() {

	fmt.Println("Starting server")
	ln, err := net.Listen("tcp", ":8080")
	checkError(err)
	var clienti clients
	// run loop forever (or until ctrl-c)
	for {
		conn, err := ln.Accept() // this blocks until connection or error
		name:=conn.RemoteAddr().String()
		clienti.name=name
		clienti.conn=conn

		checkError(err)
		go handleConnection(clienti) // a goroutine handles conn so that the loop can accept other connections

	}
}