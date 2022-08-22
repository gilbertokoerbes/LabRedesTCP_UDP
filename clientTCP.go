package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileopen, _ := os.ReadFile("./file10k.txt")
	byteArray := []byte(string(fileopen))
	fmt.Println("Transmition data length: ", len(byteArray))
	fmt.Fprintf(c, (string(byteArray))+"\n")
	message, _ := bufio.NewReader(c).ReadString('\n')
	fmt.Print("->: " + message)
	// for {

	// 	//reader := bufio.NewReader(fileopen)
	// 	fmt.Print(">> ")
	// 	//text, _ := reader.ReadString('\n')
	// 	//var text=""
	// 	fmt.Fprintf(c, text+"\n")

	// 	message, _ := bufio.NewReader(c).ReadString('\n')
	// 	fmt.Print("->: " + message)
	// 	if strings.TrimSpace(string(text)) == "STOP" {
	// 		fmt.Println("TCP client exiting...")
	// 		return
	// 	}
	// }
}
