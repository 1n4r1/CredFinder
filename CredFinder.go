package main

import (
	"flag"
	"fmt"
	"time"
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	const header = `
 ____________________________________________________________

   CredFinder - GOlang Credential Finder
 ____________________________________________________________

`
	// defaultWordlist := [3]string{"password", "id", "credential"}

	// フラグの定義
	var (
		help    bool
		version bool
	)

	path := flag.String("path", "./", "Path to start credential searching")
	dictionary := flag.String("dictionary", "Default", "Dictionary for keywork searching")
	flag.BoolVar(&help, "help", false, "Show help of the program")
	flag.BoolVar(&version, "version", false, "Show version of the program")

	// フラグの解析
	flag.Parse()

	if help {
		fmt.Println("Usage:")
		fmt.Println(".\\CredFinder.exe -Path C:\\ -Dictionary dictionary.txt")
		return
	}

	// バージョン情報のチェック
	if version {
		fmt.Println("Version 1.0.0")
		return
	}

	startTime := time.Now()

	fmt.Printf(header)
	fmt.Printf("Start searching possible credential under %s\n", *path)
	fmt.Printf("Dictionary: %s\n", *dictionary)
	fmt.Println("Started at:", startTime)
	fmt.Println("=======================Result=========================")

	searchTerm := "password"
	err := filepath.Walk(*path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			found, err := searchInFile(path, searchTerm)
			if err != nil {
				fmt.Println("Error reading file:", path, err)
				return nil
			}
			if found {
				fmt.Println("Found in:", path)
			}
		} else {
			fmt.Println("Searching folder: ", path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path", *path, err)
	}


	fmt.Println("\n=======================Finished=======================")
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("Finished at:", endTime)
	fmt.Println("Execution time:", duration)
}

func searchInFile(filePath string, term string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), term) {
			return true, nil
		}
	}
	return false, scanner.Err()
}