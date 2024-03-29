package config

type ArchitecureType string

const (
	Clean     ArchitecureType = "clean"
	Onion     ArchitecureType = "onion"
	Hexagonal ArchitecureType = "hexagonal"
	DDD       ArchitecureType = "ddd"
)

var ArchTypes = []string{string(Clean), string(Onion), string(Hexagonal), string(DDD)}
