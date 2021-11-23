/*package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type clients struct{
	name string
	conn net.Conn
}
type P struct {
	M, N int64
}
type sendcl struct{

	arr []int64
}

func handleConnection(client clients) {
	conn:=client.conn
	name:=client.name
	dec := gob.NewDecoder(conn)
	//p:=&P{}
	p := []int64{}
	dec.Decode(&p)
	fmt.Printf("Received : %+v", p);
	newmessage := strings.ToUpper("Am primit de la "+name)
	// send new string back to client
	conn.Write([]byte(newmessage + "\n"))
	time.Sleep(100*time.Millisecond)
	conn.Write([]byte("Server is processing" + "\n"))


	encoder := gob.NewEncoder(conn)
	fmt.Fprintf(conn,"Server a transmis %+v catre client \n",p)

	encoder.Encode(p)
	//v:=&sendcl{,[]int64{1,2,3}}
	/*var vec=[]int64{1,2,3}
	v:=&sendcl{[]int64{1,2,3}}
	ms:="123"
	//BinaryMarshaler{{1,2,3},e }
	fmt.Println("De trimis %+v;;;;%+v",v,vec)
	fmt.Fprintf(conn,"Server a transmit %d catre client",vec)
	encoder.Encode(&ms)
	*/
	fmt.Println("done");



	//conn.Write([]byte(string(v)))
	conn.Close()

}


func main() {

	fmt.Println("start");
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	var clienti clients
	countcl:=1
	// run loop forever (or until ctrl-c)
	for {
		conn, err := ln.Accept() // this blocks until connection or error
		name:="client"+strconv.Itoa(countcl)
//		clienti[countcl].name=name
//		clienti[countcl].conn=conn
		clienti.name=name
		clienti.conn=conn

		countcl++

		if err != nil {
			// handle error
			continue
		}
		fmt.Println(conn.RemoteAddr())
		go handleConnection(clienti) // a goroutine handles conn so that the loop can accept other connections

		//conn.Close()
	}
}