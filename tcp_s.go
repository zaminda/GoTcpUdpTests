package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// func main() {
// 	arguments := os.Args
// 	if len(arguments) == 1 {
// 		fmt.Println("Please provide port number")
// 		return
// 	}

// 	PORT := ":" + arguments[1]
// 	l, err := net.Listen("tcp", PORT)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer l.Close()

// 	c, err := l.Accept()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	for {
// 		netData, err := bufio.NewReader(c).ReadString('\n')
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		if strings.TrimSpace(string(netData)) == "STOP" {
// 			fmt.Println("Exiting TCP server!")
// 			return
// 		}

// 		fmt.Print("-> ", string(netData))
// 		t := time.Now()
// 		myTime := t.Format(time.RFC3339) + "\n"
// 		c.Write([]byte(myTime))
// 	}
// }

// import (
//     "bufio"
//     "fmt"
//     "log"
//     "net"
//     "strings"
// )

func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	//NewReader(c).ReadString('\n')
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println("Message Received:", message)
		//newMessage := strings.ToUpper(message)
		//conn.Write([]byte(newMessage + "\n"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
}

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	socket := "127.0.0.1:" + arguments[1]

	ln, err := net.Listen("tcp", socket)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Accept connection on port")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Calling handleConnection")
		go handleConnection(conn)
	}
}
