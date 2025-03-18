package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("err opening: ", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 8)

	for {
		thing, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err opening: ", err)
			return
		}

		fmt.Printf("read: %s\n", buffer[:thing])
	}
}
