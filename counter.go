package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Count(url string) int {
	response, err := http.Get(url)
	if err != nil {
		log.Println("error!!!", err)
		return 0
	}

	if response.StatusCode != 200 {
		log.Println("status.Code", response.StatusCode)
		return 0
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Println("error while closing body", err)
		}
	}()

	body, err := ioutil.ReadAll(response.Body)
	numberGo := strings.Count(string(body), "Go")
	log.Printf("count for url %v: %v", url, numberGo)
	return numberGo
}
