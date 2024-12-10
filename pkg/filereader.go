package pkg

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func mustOpenFile(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

// DirFile returns the absolute path of the chosen input file in the current directory.
func DirFile(name string) string {
	_, p, _, _ := runtime.Caller(1)
	return filepath.Join(filepath.Dir(p), "../inputs/", name)
}

func MustReadFile(filename string) string {
	f := mustOpenFile(DirFile(filename))
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
