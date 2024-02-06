package main

import (
	"fmt"
)

func main() {
	mapf, reducef := loadPlugin()
	fmt.Printf("Map function return value: %v\n", mapf)
	fmt.Printf("Reduce function return value: %v", reducef)

}

func loadPlugin() (int, int) {
	return 0, 0
}
