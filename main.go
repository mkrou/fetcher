package main

import (
	"github.com/mkrou/fetcher/fetcher"
	"log"
	"sync"
)

func main() {
	mock := fetcher.NewFetcher(2)

	wait := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wait.Add(2)

		go func() {
			printFetchedData(mock.Get())
			wait.Done()
		}()

		go func() {
			printFetchedData(mock.List())
			wait.Done()
		}()
	}

	wait.Wait()
}

func printFetchedData(data interface{}, err error) {
	if err != nil {
		log.Fatal("Fetching error: ", err.Error())
	}

	log.Print(data)
}
