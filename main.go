package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	// * Max Buffer size for the file
	maxSize = 1024 * 1024
)

var (
	// * Keywords and file formats to search for
	keywords           = []string{"github", ".com", "//", "/*", "*/"}
	file_formats       = []string{".txt", ".md", ".json", ".yaml", ".yml", ".ts", ".js", ".sh", ".go", ".py", ".html", ".css", ".cs", ".java", ".c", ".o", ".h"}
	restricted_folders = []string{"node_modules", ".vscode", ".git"}
	file_name          = "results.txt"
)

func main() {
	dir := os.Args[1]
	log.Println("Searching dir: ", dir)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalln(err)
		}

		if info.IsDir() {
			// * Check if the dir is allowed to be searched
			for _, restricted_dir := range restricted_folders {
				if restricted_dir == info.Name() {
					log.Println("Folder is restricted for searching")
					return filepath.SkipDir
				}
			}
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
		key := "/;#;/"
		// * Check the line for keywords
		for _, i := range keywords {
			if strings.Contains(line, i) {
				writeToFile(line + key + path)
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func writeToFile(data string) {
	f, err := os.OpenFile(file_name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(data + "\n")); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
