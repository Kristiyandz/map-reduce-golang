package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"plugin"

	"example.com/models"
)

func main() {

	mapf := loadPlugin(os.Args[1])

	fmt.Println(os.Args[2:])

	intermidiate := []models.KeyValue{}

	// can take many file names and process them in a loop
	// an improvement will be to send each file into a goroutine
	for _, filename := range os.Args[2:] {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("cannot open %v", file)
		}

		// whole file contents
		content, err := io.ReadAll(file)
		if err != nil {
			log.Fatalf("cannot read content %v", content)
		}
		file.Close()
		// pass the file name and the content
		kva := mapf(filename, string(content)) // Pass the content as the second argument
		intermidiate = append(intermidiate, kva...)

	}
	// fmt.Printf("The intermidiate: %v", intermidiate)

}

func loadPlugin(fileName string) func(string, string) []models.KeyValue {
	// make sure this is ran: go build -buildmode=plugin + path name
	file, err := plugin.Open(fileName)
	if err != nil {
		log.Fatalf("Cannot open plugin %v", fileName)
	}

	mapFuncImported, err := file.Lookup("Map")
	if err != nil {
		log.Fatalf("Cannot find map in file %v", fileName)
	}

	// go run main.go ../mapreduce.so ../mapreduce/lotr-script.txt
	mapFunc, ok := mapFuncImported.(func(string, string) []models.KeyValue)

	if !ok {
		log.Fatalf("Invalid type for map function in file %v", fileName)
	}

	return mapFunc

}
