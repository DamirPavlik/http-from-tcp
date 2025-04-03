package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const inputFilePath = "messages.txt"

const port = ":42069"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("could not listen: ", err)
	}
	defer lis.Close()

	fmt.Println("listening for tcp traffic on: ", port)
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("error accepting: ", err)
		}
		fmt.Println("accepted conn from: ", conn.RemoteAddr())
		linesChan := getLinesChannel(conn)

		for line := range linesChan {
			fmt.Println(line)
		}
		fmt.Println("Connection to ", conn.RemoteAddr(), "closed")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)
	go func() {
		defer f.Close()
		defer close(lines)
		currentLineContents := ""
		for {
			b := make([]byte, 8, 8)
			n, err := f.Read(b)
			if err != nil {
				if currentLineContents != "" {
					lines <- currentLineContents
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				return
			}
			str := string(b[:n])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts)-1; i++ {
				lines <- fmt.Sprintf("%s%s", currentLineContents, parts[i])
				currentLineContents = ""
			}
			currentLineContents += parts[len(parts)-1]
		}
	}()
	return lines
}
