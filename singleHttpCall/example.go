package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type AsyncResult struct {
	Response string
	Error    error
}

func makeHttpCallWithChannel(ch chan<- *AsyncResult) {
	resp, err := makeHttpCall()
	ch <- &AsyncResult{
		Response: resp,
		Error:    err,
	}
}

func makeHttpCall() (string, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	// Example using a wait group.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		responseString, err := makeHttpCall()
		if err != nil {
			log.Fatal(err)
		}

		log.Println(responseString)
	}()

	wg.Wait()

	// Example using channels.
	ch := make(chan *AsyncResult)
	go makeHttpCallWithChannel(ch)
	result := <-ch
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println(result.Response)
}
