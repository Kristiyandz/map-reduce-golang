package main

import (
	"fmt"
	"log"
	"os"
	"plugin"
)

// this is a duplicate struct, figure out how to use just a single one
// type KeyValue struct {
// 	Key   string
// 	Value string
// }

func main() {
	fmt.Println(os.Args[1])
	mapf := loadPlugin(os.Args[1])
	fmt.Printf("Map function return value: %v\n", mapf)
	// fmt.Printf("Reduce function return value: %v", reducef)

}

func loadPlugin(fileName string) int {
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

	mapFunc := mapFuncImported

	fmt.Printf("map func %v", mapFunc)

	// return mapFunc
	return 0

}
