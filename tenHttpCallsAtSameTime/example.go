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
	results := make([]AsyncResult, 10)
	for i, _ := range results {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			resp, err := makeHttpCall()
			results[i] = AsyncResult{
				Response: resp,
				Error:    err,
			}
		}(i)
	}

	wg.Wait()

	for _, result := range results {
		if result.Error != nil {
			log.Fatal(result.Error)
		}

		log.Println(result.Response)
	}
}
