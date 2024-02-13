package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"plugin"
	"reflect"

	"example.com/models"
)

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

func loadPlugin(fileName string) func(string, string) []models.KeyValue {
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

	// go run main.go ../mapreduce.so ../mapreduce/lotr-script.txt
	// type of imported func is wrong: func(string, string) []main.KeyValue
	mapFunc, ok := mapFuncImported.(func(string, string) []models.KeyValue)
	fmt.Println("but not here")
	fmt.Printf("hubabuba %v", ok)
	if !ok {
		log.Println(reflect.TypeOf(mapFuncImported))

		log.Fatalf("Invalid type for map function in file %v", fileName)
	}

	return mapFunc

}
