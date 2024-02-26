package utils

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

// Colorise text in the terminal
func Colorise(color string, text string) string {
	colors := map[string]string{
		"reset":   "\033[0m",
		"black":   "\033[30m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
	}
	return colors[color] + text + colors["reset"]
}

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

func GetField(p interface{}, fieldName string) *string {
	v := reflect.ValueOf(p)
	f := reflect.Indirect(v).FieldByName(fieldName)
	return f.Addr().Interface().(*string)
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

func CLIParser(p interface{}) *flag.FlagSet {
	fs := flag.NewFlagSet("parser", flag.ExitOnError)

	t := reflect.TypeOf(p)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// get the tags of the field
		tag := field.Tag
		short := tag.Get("short")
		long := tag.Get("long")
		description := tag.Get("description")

		// ensure that the short and long flags are not empty
		if short == "" || long == "" {
			continue
		}

		fmt.Println(field)

		// Add the field to the flagset
		fs.StringVar(GetField(p, field.Name), short, "", description)
		fs.StringVar(GetField(p, field.Name), long, "", description)
	}

	return fs
}

// Loading represents a loading animation
type Loading struct {
	quit chan struct{}
}

// Start starts the loading animation with the given message
func (l *Loading) Start(message string) {
	fmt.Print(Colorise("green", message))
	l.quit = make(chan struct{})

	go func() {
		for {
			select {
			case <-l.quit:
				return
			default:
				fmt.Print("\b/") // Spinner position 1
				time.Sleep(100 * time.Millisecond)
				fmt.Print("\b-") // Spinner position 2
				time.Sleep(100 * time.Millisecond)
				fmt.Print("\b\\") // Spinner position 3
				time.Sleep(100 * time.Millisecond)
				fmt.Print("\b|") // Spinner position 4
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}

// Stop stops the loading animation
func (l *Loading) Stop() {
	close(l.quit)
	fmt.Println(" Done")
}

// Fetch username from git config
func GetGitUsername() string {
	cmd := exec.Command("git", "config", "user.name")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(out))
}

func OutputPathHandler(outputPath string) string {
	// check if the output path has ./ or ../
	if strings.HasPrefix(outputPath, "./") || strings.HasPrefix(outputPath, "../") {
		// get the current working directory
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		// join the current working directory with the output path
		absPath := filepath.Join(cwd, outputPath)
		return absPath
	}

	// if it doesn't then add ./ to the output path
	return filepath.Join(".", outputPath)

}
