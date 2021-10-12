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
	// * Buffer size for the file
	maxSize = 1024 * 1024
)

var (
	keywords     = []string{"github", ".com", "//", "/*", "*/"}
	file_formats = []string{".txt", ".gitignore", ".ts", ".js", ".sum", ".mod", ".md", ".sh", ".json", ".yaml", ".lock", ".tf", ".go", ".py", ".groovy", ".csh", ".html", ".css"}
	file_name    = "results.txt"
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
			log.Fatalln(err)
		}
		if !info.IsDir() {
			// * If the file isn't diretory and matches the file format read the file
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
				writeToFile(line + "/;#;/" + path)
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func writeToFile(data string) {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(file_name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(data + "\n")); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
