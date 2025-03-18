package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("err opening: ", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 8)
	var currentLine string

	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err opening: ", err)
			return
		}

		chunk := string(buffer[:bytesRead])
		currentLine += chunk
		parts := strings.Split(currentLine, "\n")

		for i := 0; i < len(parts)-1; i++ {
			fmt.Printf("read: %s\n", parts[i])
		}

		currentLine = parts[len(parts)-1]
	}

	if len(currentLine) > 0 {
		fmt.Printf("read: %s\n", currentLine)
	}
}
