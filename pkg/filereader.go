package pkg

import (
	"io"
	"log"
	"os"
	"strings"
)

func mustOpenFile(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func MustReadFile(filename string) string {
	f := mustOpenFile(filename)
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func MustReadFileLines(filename string) []string {
	return strings.Split(MustReadFile(filename), "\n")
}
