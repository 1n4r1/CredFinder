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
	defaultWordlist := []string{"password", "id", "credential","パスワード","認証情報"}

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
	if *dictionary != "Default" {
		wordlist = defaultWordlist
	} else {
		{} // read the file and append in the array
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

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			fmt.Println("Searching folder:", path)
			return nil
		}

		found, err := searchInFile(path, wordlist)
		if err != nil {
			fmt.Println("Error reading file:", path, err)
			return nil
		}

		if found {
				fmt.Println("Found in:", path)
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

func searchInFile(path string, terms []string) (bool, error) {
	file, err := os.Open(path)
	items := []string{}
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for _, term := range terms {
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), term) {
				items = append(items, scanner.Text())
			}
		}
	}

	if items != nil {
		// To return the content of "items"
		return true, nil
	} else {
		return false, scanner.Err()
	}
}