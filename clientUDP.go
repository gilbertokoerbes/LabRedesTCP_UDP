package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	p := make([]byte, 2048)
	conn, err := net.Dial("udp", "159.223.156.81:1234")
	//conn, err := net.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fileopen, _ := os.ReadFile("./file1500.txt")
	byteArray := []byte(string(fileopen))
	fmt.Println("Transmition data length: ", len(byteArray))
	fmt.Fprintf(conn, (string(byteArray))+"\n")
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}
