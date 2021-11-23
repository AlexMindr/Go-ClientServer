package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type i struct{
	val int64
}
func chkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func arr(n int)[]int64 {
	a := make([]int64, n)
	fmt.Println("Enter the inputs")
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&a[i])
		chkError(err)
	}
	return a
}
func subiect3(conn net.Conn, nume string){

	//	p:=[]int64{12,23}
	var n int
	fmt.Println("Enter the number of elements")
	_, err14 := fmt.Scan(&n)
	chkError(err14)
	p:=arr(n)
	encoder := gob.NewEncoder(conn)

	err11 := encoder.Encode(p)
	chkError(err11)


	for i:=1;i<5;i++{
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if message != "" {
			fmt.Print("Message from server: " + message)
		}
		if strings.Contains(message,"Server sends"){
			re := regexp.MustCompile("[0-9]+")
			//fmt.Println(re.FindAllString(message, 1))
			res:=re.FindAllString(message, 1)
			rez,_:=strconv.Atoi(res[0])
			_, err10 := fmt.Fprintf(conn,"Client %s recieved %v as result\n",nume, rez)
			chkError(err10)

		}
	}


}

func main() {

	fmt.Println("Starting client")
	//nume:="Andrei"
	// connect to this socket
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Connection error", err)
	}


	var nume string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your client name: ")
	scanner.Scan()
	// Holds the string that scanned
	nume = scanner.Text()
	if len(nume) != 0 {
	} else {
		panic("Empty input\n")
	}
	//fmt.Println(nume)

	time.Sleep(1000*time.Millisecond)
	_, err13 := fmt.Fprintf(conn,"%s\n",nume)
	chkError(err13)

	var subiect string
	fmt.Println("Introduceti subiectul de rezolvat:")
	scanner.Scan()
	// Holds the string that scanned
	subiect = scanner.Text()
	if len(subiect) != 0 {
	} else {
		panic("Empty input\n")
	}

	_, err14 := fmt.Fprintf(conn,"%s\n",subiect)
	chkError(err14)

	switch subiect {
	case "3":
		subiect3(conn,nume)
	default:
		fmt.Println("Nu exista subiectul")
	}
	err8 := conn.Close()
	chkError(err8)

}