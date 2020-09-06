package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func FormatGoFilesInDir(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		if info.Mode().IsRegular() {
			if strings.Contains(path, ".go") {
				Format(path)
			}
		}
		return nil
	})
}

func Format(filePath string) {
	cmd := exec.Command("", "")

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "gofmt", "-w", filePath)
	} else {
		cmd = exec.Command("gofmt", "-w", filePath)
	}
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "*** fmt error ***", err)
	} else {
		fmt.Printf("%s\n", filePath)
	}
}

func main() {
	var pathParam string
	flag.StringVar(&pathParam, "path", ".", "Scan path for go files")
	flag.Parse()

	FormatGoFilesInDir(pathParam)
}
