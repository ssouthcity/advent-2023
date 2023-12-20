package main

import (
	"os"

	"github.com/BurntSushi/toml"
)

type DayMeta struct {
	Description string
	Song        struct {
		Name     string
		Features []string
	}
}

type Metadata struct {
	Days map[string]DayMeta
}

func parseMetaFile(path string) (Metadata, error) {
	metadata := Metadata{}

	f, err := os.Open(path)
	if err != nil {
		return metadata, err
	}
	defer f.Close()

	d := toml.NewDecoder(f)

	_, err = d.Decode(&metadata)
	if err != nil {
		return metadata, err
	}

	return metadata, nil
}
