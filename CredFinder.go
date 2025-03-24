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
	defaultWordlist := []string{"password", "id", "credential"}

	var (
		help    bool
		version bool
	)

	root := flag.String("path", "./", "Path to start credential searching")
	dictionary := flag.String("dictionary", "Default", "Dictionary for keyword searching")
	//host := flag.String("host", "localhost", "Hostname to run the script")
	//hosts := flag.String("hosts", "hosts.txt", "Host list to run the script")
	flag.BoolVar(&help, "help", false, "Show help of the program")
	flag.BoolVar(&version, "version", false, "Show version of the program")

	flag.Parse()

	if help {
		fmt.Println("Usage:")
		fmt.Println(".\\CredFinder.exe -Path C:\\ -Dictionary dictionary.txt")
		return
	}

	wordlist := defaultWordlist	
	if *dictionary == "Default" {
		wordlist = defaultWordlist
	} else {
		file, err := os.Open(*dictionary)
		if err != nil {
			fmt.Println("Error dictionary file:", *dictionary, err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			wordlist = append(wordlist, scanner.Text())
		}
	}

	if version {
		fmt.Println("Version 1.0.0")
		return
	}

	startTime := time.Now()
	rootDir := *root

	fmt.Printf(header)
	fmt.Printf("Start searching possible credential under \"%s\"\n", rootDir)
	fmt.Printf("Dictionary: %s\n", wordlist)
	fmt.Println("Started at:", startTime.Format("2006-01-02 15:04:05"))
	fmt.Println("=======================Result=========================")

	// multibyte characters are dealt at last, wanna quit using filepath.Walk()
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, er error) error {
		if er != nil {
			return er
		}

		if info.IsDir() {
			// fmt.Println("Searching folder:", path)
			return nil
		}

		searchFileName(path, info.Name() , wordlist)
		e := searchInFile(path, wordlist)
		if e != nil {
			fmt.Println("Error reading file:", path, e)
			return nil
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking the root folder", rootDir, err)
	}

	fmt.Println("\n=======================Finished=======================")
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("Finished at:", endTime.Format("2006-01-02 15:04:05"))
	fmt.Println("Execution time:", duration)
}

func searchFileName(path string, filename string, terms []string) {
	for _, term := range terms{
		if strings.Contains(filename, term) {
			fmt.Printf("Interesting filename: \"%s\"\n", path)
		}
	}
}

func searchInFile(path string, terms []string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, term := range terms {
			if strings.Contains(scanner.Text(), term) {
				fmt.Printf("Found \"%s\" in: %s\n", term, path)
			}
		}
	}
	return scanner.Err()
}