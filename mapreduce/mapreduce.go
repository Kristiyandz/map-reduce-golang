package mapreduce

import (
	"strings"
	"unicode"

	"example.com/models"
)

func Map(fileName string, contents string) []models.KeyValue {
	ff := func(r rune) bool { return !unicode.IsLetter(r) }

	words := strings.FieldsFunc(contents, ff)

	keyValueSlice := []models.KeyValue{}

	for _, w := range words {
		kv := models.KeyValue{Key: w, Value: "1"}
		keyValueSlice = append(keyValueSlice, kv)
	}
	return keyValueSlice
}
