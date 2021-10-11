package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	maxSize = 1024 * 1024
)

var (
	results      = []string{}
	keywords     = []string{"github", "aws", "//", "/*", "*/"}
	file_formats = []string{".txt", ".gitignore", ".ts", ".js", ".sum", ".mod", ".md", ".sh", ".json", ".yaml", ".lock", ".tf", ".go", ".py", ".groovy", ".csh", ".html", ".css"}
)

func main() {
	log.Println("Starting Code Searcher...")
	// * Variables
	var dir string
	// * Read the flag
	flag.StringVar(&dir, "d", ".", "Specifies the directory to search. Default searches the current dir")
	flag.Parse()
	// * Directory looping
	log.Println("Searching dir: ", dir)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
		}
		if !info.IsDir() {
			// * Add the task to the wait queue and read file
			for _, ext := range file_formats {
				if filepath.Ext(path) == ext {
					ReadFile(path)
				}
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	writeToFile()
}

func ReadFile(path string) {
	log.Println("Searching in: ", path)
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	buffer := make([]byte, 0, maxSize)
	scanner.Buffer(buffer, maxSize)
	for scanner.Scan() {
		line := string(scanner.Text())
		// * Check the line for keywords
		for _, i := range keywords {
			if strings.Contains(line, i) {
				results = append(results, line)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func writeToFile() {
	file, err := os.Create("result.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for _, line := range results {
		// * Loop through the results in results and write to file
		file.WriteString(line + "\n")
	}
}
