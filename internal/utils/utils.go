package utils

import (
	"fmt"
	"reflect"
	"strings"
)

// NormalizeString normalizes a string by trimming spaces and converting it to lowercase
func NormalizeString(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}

// PrintError prints an error message
func PrintError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// GetFields returns the fields of a struct
func GetFields(p interface{}) (fields []string) {
	fields = make([]string, 0)

	// Dereference the pointer if p is a pointer
	t := reflect.TypeOf(p)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Check if the type is a struct
	if t.Kind() != reflect.Struct {
		fmt.Println("Not a struct type")
		return fields
	}

	// Parse the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// Do something with the fields
		fields = append(fields, field.Name)
	}

	return fields
}

// GetTags returns the tags of a struct
func GetTags(p interface{}, tags ...string) (tagValues map[string][]string) {

	tagValues = make(map[string][]string)

	// Parse the tags of the struct
	t := reflect.TypeOf(p)

	// Dereference the pointer if p is a pointer
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Check if the type is a struct
	if t.Kind() != reflect.Struct {
		fmt.Println("Not a struct type")
		return tagValues
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag
		for _, tagKey := range tags {
			tagValue := tag.Get(tagKey)
			tagValues[tagKey] = append(tagValues[tagKey], tagValue)
		}
	}

	return tagValues

}
