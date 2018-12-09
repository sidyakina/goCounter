package main

import (
	"bufio"
	"log"
	"strings"
)

func init() {
	log.SetFlags(0)
}

func main() {
	// generate string with url
	numberUrls := 11
	urls := "https://golang.org"
	for i := 0; i < numberUrls-1; i++ {
		urls += "\nhttps://golang.org"
	}
	r := strings.NewReader(urls)

	maxRoutines := 3
	scanner := bufio.NewScanner(r)
	tasker := InitTasker(maxRoutines)

	for scanner.Scan() {
		url := strings.Trim(scanner.Text(), " ")
		tasker.Add(url)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	total := tasker.Close()
	log.Printf("Total: %v", total)
}
