package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"plugin"
)

// this is a duplicate struct, figure out how to use just a single one
type KeyValue struct {
	Key   string
	Value string
}

func main() {
	fmt.Println(os.Args[1])
	fmt.Println("here")
	mapf := loadPlugin(os.Args[1])

	fmt.Printf("%v", mapf)

	// for _, filename := range os.Args[2:] {
	file, err := os.Open("./lotr-script.txt")
	if err != nil {
		log.Fatalf("cannot open %v", file)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("cannot read content %v", content)
	}
	file.Close()
	// kva := mapf(filename, string(1)) // Pass the content as the second argument
	// fmt.Printf("kva: %v", kva)
	// }
	fmt.Printf("Map function return value: %v\n", mapf)
	// fmt.Printf("Reduce function return value: %v", reducef)

}

func loadPlugin(fileName string) func(string, string) []KeyValue {
	/*
		Input data will be read here;
		The data will be passed into the Map function
		The intermidiate data produced from Map will be accumulated and returned
	*/

	file, err := plugin.Open(fileName) // to access this, it needs to be built as plugin: go build -buildmode=plugin + path name

	if err != nil {
		log.Fatalf("Cannot open plugin %v", fileName)
	}

	mapFuncImported, err := file.Lookup("Map")
	if err != nil {
		log.Fatalf("Cannot find map in file %v", fileName)
	}

	fmt.Println("there")

	// type is fuqed
	mapFunc, ok := mapFuncImported.(func(string, string) []KeyValue)
	fmt.Println("but not here")
	fmt.Printf("hubabuba %v", ok)
	if !ok {
		log.Fatalf("Invalid type for map function in file %v", fileName)
	}

	return mapFunc

}
