package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputFilePath = "messages.txt"

func main() {
	f, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("could not open %s: %s\n", inputFilePath, err)
	}

	fmt.Printf("reading data from %s\n", inputFilePath)

	currentLineContent := ""
	for {
		buffer := make([]byte, 8, 8)
		n, err := f.Read(buffer)

		if err != nil {
			if currentLineContent != "" {
				fmt.Printf("read: %s\n", currentLineContent)
				currentLineContent = ""
			}
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("error: %s\n", err.Error())
			break
		}
		str := string(buffer[:n])
		parts := strings.Split(str, "\n")
		for i := 0; i < len(parts)-1; i++ {
			fmt.Printf("read: %s%s\n", currentLineContent, parts[i])
			currentLineContent = ""
		}
		currentLineContent += parts[len(parts)-1]
	}
}
