package main

import (
	"bufio"
	"encoding/gob"
	"log"
	"net"
)
import "fmt"

type Q struct {
	M, N int64
}

type recvcl struct{

	arr []int64
}


func main() {
	fmt.Println("start client");

	// connect to this socket
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Connection error", err)
	}
	encoder := gob.NewEncoder(conn)
	//p:=&Q{1,2}
	p:=[]int64{1,2}
	encoder.Encode(p)
	fmt.Println("done");
	counter:=0

	// read in input from stdin
/*		reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, _ := reader.ReadString('\n')
	// send to socket
	fmt.Fprintf(conn, text + "\n")
*/		// listen for reply
	//de 5 ori si apoi trimite
	for i:=1;i<=3;i++{
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if (message != "") {
			fmt.Print("Message from server: " + message)
		}

	}
	decsv := gob.NewDecoder(conn)
	//resv := &[]int64{}
	//re:=&Q{}
	re:=[]int64{}

	decsv.Decode(&re)
	fmt.Printf("Received : %+v", re);

	//conn.Write()
	conn.Close()
	counter++

}