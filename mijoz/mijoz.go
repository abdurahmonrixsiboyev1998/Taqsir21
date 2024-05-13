package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

func Client() {
	var wg sync.WaitGroup
	files := os.Args[1:]

	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			sendFile(file)
		}(file)
	}

	wg.Wait()
}

func sendFile(filePath string) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = conn.Write(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Fayl yuborildi:", filePath)
}

func main() {
	fileName := "./fileName.txt"
	sendFile(fileName)
}
