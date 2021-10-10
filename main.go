package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	// * keywords to search for in file
	keywords = []string{"tv", "github", "//", "aws"}
	results  = make(chan string)
)

func main() {
	log.Println("Starting Code Searcher...")

	// * Variables
	var dir string
	var wg sync.WaitGroup

	// * Read the flag
	flag.StringVar(&dir, "d", ".", "Specifies the directory to search. Default searches the current dir")
	flag.Parse()

	// * Directory looping
	log.Println("Searching dir: ", dir)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
			return nil
		}
		if !info.IsDir() {
			// * Add the task to the wait queue and read file
			wg.Add(1)
			go func() {
				defer wg.Done()
				ReadFile(path)
			}()
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	go func() {
		// * When the wait queue ends close the channel
		wg.Wait()
		close(results)
	}()
	for i := range results {
		// * Loop throught the results in channel and print it
		log.Println(i)
	}
}

func ReadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := string(scanner.Text())
		// * Check the line for keywords
		for _, i := range keywords {
			if strings.Contains(line, i) {
				results <- line
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
