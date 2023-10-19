package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d bytes\n", n2)

	data, err := os.ReadFile("test2.txt")
	if err != nil {
		panic(err)
	}
	for _, v := range data {
		fmt.Printf("wow %v\n", v)
	}
}
