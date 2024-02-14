package main

import (
	"strings"
	"unicode"

	"example.com/models"
)

func Map(fileName string, contents string) []models.KeyValue {
	// remove all symbols such as comma and periods and white space
	ff := func(r rune) bool { return !unicode.IsLetter(r) }
	words := strings.FieldsFunc(contents, ff)

	keyValueSlice := []models.KeyValue{}
	for _, w := range words {
		// this is 1 so it refers to the first files passed
		// it will increment as the number of files increases
		kv := models.KeyValue{Key: w, Value: "1"}
		keyValueSlice = append(keyValueSlice, kv)
	}
	return keyValueSlice
}
