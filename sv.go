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



var countinst=0

func citireLinii(cale string) ([]string, error) {
	fisier, err := os.Open(cale)
	if err != nil {
		return nil, err
	}

	//sa astepte ca functiile de mai sus (portiunea de cod) sa se execute corespunzator
	//inchidem fisierul
	defer func(fisier *os.File) {
		err := fisier.Close()
		checkError(err)
	}(fisier)

	var linii []string
	scanner := bufio.NewScanner(fisier)
	for scanner.Scan() {
		linii = append(linii, scanner.Text())
	}
	return linii, scanner.Err()
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

func convert(s string) int64{
	var res int64
	res, _ = strconv.ParseInt(s, 2, 64)
	//fmt.Println(res)
	return res
}
func strtoint(s []string, c chan []int64) {

	var res []int64

	for _, v := range s {
		re := regexp.MustCompile("^[01]+$")
		//fmt.Println(re.FindAllString(message, 1))
		ok:=re.MatchString(v)
		if ok == true{
			res=append(res,convert(v))

		}
	}
	fmt.Println(res)
	c <- res // send sum to c
}

func sub5(conn net.Conn,name string){
	dec := gob.NewDecoder(conn)
	var p []string
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

	c := make(chan []int64)
	go strtoint(p, c)
	btoint:=<-c


	_, err4 := fmt.Fprintf(conn,"Server sends %d as reply \n",btoint)
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
	case "5":
		sub5(conn,name)
	default:
		fmt.Println("Nu exista subiectul!")
	}
	fmt.Printf("Closing connection\n")
	countinst=countinst-1
}


func main() {

	//deschide fisierul pentru citire, linie cu linie
	linii, errf := citireLinii("config.txt")
	if errf != nil {
		log.Fatalf("citireLinii: %s", errf)
	}
	var file []string
	// afisam continutul fisierului
	for _, linie := range linii {

		file=append(file,linie)
	}
	port:=file[0]
	var nrclienti int
	nrclienti, _ =strconv.Atoi(file[1])
	fmt.Println("Starting server on port "+port +" max number of clients allowed: " + strconv.Itoa(nrclienti))
	ln, err := net.Listen("tcp", port)
	checkError(err)
	var clienti clients

	// run loop forever (or until ctrl-c)
	for {
		conn, err := ln.Accept() // this blocks until connection or error
		name:=conn.RemoteAddr().String()
		clienti.name=name
		clienti.conn=conn

		checkError(err)
		if countinst<nrclienti {
			_, err := fmt.Fprintf(conn,"Connected\n")
			checkError(err)
			time.Sleep(5000*time.Millisecond)
			go handleConnection(clienti) // a goroutine handles conn so that the loop can accept other connections
			countinst++
		}else {
			_, err := fmt.Fprintf(conn,"Disconnected,server is full\n")
			checkError(err)
			fmt.Println("Please upgrade the server!")
		}
	}
}