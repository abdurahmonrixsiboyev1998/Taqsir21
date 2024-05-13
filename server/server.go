package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listens, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listens.Close()
	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listens.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(conn, "./New.txt")
	}
}

func handleConnection(conn net.Conn, fileName string) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	filestr := string(buf[:n])

	newFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer newFile.Close()

	_, err = newFile.WriteString(filestr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Fayl yuklandi:", fileName)
}
