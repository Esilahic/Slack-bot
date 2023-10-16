package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	txt := []byte("this is a test\n")
	err := os.WriteFile("test.txt", txt, 0644)
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create("test2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	d2 := []byte{1, 2, 3}
	_, err = f.Write(d2)
	if err != nil {
		log.Fatal(err)
	}
}
