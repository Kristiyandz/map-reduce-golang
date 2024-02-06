package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	mapf, reducef := loadPlugin()
	fmt.Printf("Map function return value: %v\n", mapf)
	fmt.Printf("Reduce function return value: %v", reducef)

}

func loadPlugin() (int, int) {
	/*
		Input data will be read here;
		The data will be passed into the Map function
		The intermidiate data produced from Map will be accumulated and returned
	*/

	for _, fileName := range os.Args[1:] {
		file, err := os.Open(fileName)

		content, err := io.ReadAll(file)

		if err != nil {
			log.Fatalf("cannot read %v", fileName)
		}
		file.Close()

		// pass the file name and the content to the Map
		fmt.Println(content)
	}
	return 0, 0
}
