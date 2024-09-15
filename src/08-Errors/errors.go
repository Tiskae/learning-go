package main

import (
	"errors"
	"fmt"

	// "log"
	"os"
	// "panic"
	"strings"
)

func isInvalidGithubLink(link string) error {
	if !strings.Contains(link, "github.com/") {
		err := errors.New("invalid github link")
		return err
	}

	return nil
}

func openFile(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func main() {
	githubLink := "https://github.coom/tiskae"
	if err := isInvalidGithubLink((githubLink)); err != nil {
		fmt.Println(err)
		// panic(err)
		// log.Fatal(err)
	}

	if err := openFile("missing-file.txt"); err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}
}
