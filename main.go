package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const notRegularFileMode os.FileMode = os.ModeDir | os.ModeSymlink | os.ModeNamedPipe | os.ModeSocket | os.ModeDevice | os.ModeCharDevice | os.ModeIrregular

var sizes = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}

type DirFlags []string

func (i *DirFlags) String() string {
	return "Directory name"
}

func (i *DirFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func formatSize(s float64, base float64) string {
	unitsLimit := len(sizes)
	i := 0
	for s >= base && i < unitsLimit {
		s = s / base
		i++
	}

	f := "%.0f %s"
	if i > 1 {
		f = "%.2f %s"
	}

	return fmt.Sprintf(f, s, sizes[i])
}

func GetDirectorySize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if info != nil && (info.Mode()&notRegularFileMode) == 0 {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func GetCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	return dir
}

func Run(path string) {
	size, err := GetDirectorySize(path)
	if err != nil {
		fmt.Println("Directory or file not found.")
		os.Exit(-1)
	}
	fmt.Printf("%s: %s\n", path, formatSize(float64(size), 1024.0))
}

func main() {
	var Dirs = DirFlags{GetCurrentDir()}
	flag.Var(&Dirs, "dir", "Directory name")
	flag.Usage = func() {
		fmt.Printf("Usage: dirsize [OPTIONS] argument ...\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(os.Args) >= 2 && !strings.HasPrefix(os.Args[1], "--") {
		Dirs = append(Dirs, os.Args[1:]...)
	}

	dirLen := len(Dirs)
	if dirLen >= 2 {
		Dirs = Dirs[1:]
	}

	for _, path := range Dirs {
		Run(path)
	}
}
