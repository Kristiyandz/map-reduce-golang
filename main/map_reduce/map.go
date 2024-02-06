package main

import (
	"strings"
	"unicode"
)

type KeyValue struct {
	Key   string
	Value string
}

func Map(fileName string, contents string) []KeyValue {
	ff := func(r rune) bool { return !unicode.IsLetter(r) }

	words := strings.FieldsFunc(contents, ff)

	keyValueSlice := []KeyValue{}

	for _, w := range words {
		kv := KeyValue{w, "1"}
		keyValueSlice = append(keyValueSlice, kv)
	}
	return keyValueSlice
}
