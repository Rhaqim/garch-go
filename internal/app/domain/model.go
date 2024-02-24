package domain

import (
	"reflect"
)

// ProjectConfig represents the configuration for a project
type ProjectConfig struct {
	Arch   string `short:"a" long:"arch" description:"Architecture type"`
	Title  string `short:"t" long:"title" description:"Title of the project"`
	Author string `short:"u" long:"author" description:"Author of the project"`
	DbType string `short:"d" long:"db" description:"Database type"`
	// Add more fields for other configurations
}

func (p *ProjectConfig) ParseTags() (short, long, desc []string) {

	short = make([]string, 0)
	long = make([]string, 0)
	desc = make([]string, 0)

	// Parse the tags of the struct
	t := reflect.TypeOf(*p)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag
		shorts := tag.Get("short")
		longs := tag.Get("long")
		descs := tag.Get("description")
		// Do something with the tags

		short = append(short, shorts)
		long = append(long, longs)
		desc = append(desc, descs)

	}

	return short, long, desc

}

func (p *ProjectConfig) GetFields() (fields []string) {

	fields = make([]string, 0)

	// Parse the tags of the struct
	t := reflect.TypeOf(*p)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// Do something with the tags
		fields = append(fields, field.Name)

	}

	return fields

}
