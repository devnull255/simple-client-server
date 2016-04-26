package main

import "net"
import "os"
import "fmt"
import "bufio"
import "strings" // only needed below for sample processing

func handleRequest(conn net.Conn) {
	message, _ := bufio.NewReader(conn).ReadString('\n')
	// output message received
	if strings.Contains(string(message), "quit") {
		fmt.Print("quit command received.")
		os.Exit(0)
	}
	fmt.Print("Message Received:", string(message))
	// sample process for string received
	newmessage := strings.ToUpper(message)
	// send new string back to client
	conn.Write([]byte(newmessage + "\n"))
	conn.Close()

}
func main() {
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// run loop forever (or until crtl-c)
	for {
		// will listen for message to process ending in newline (\n)
		// accept connection on port
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine
		go handleRequest(conn)
	}
}
