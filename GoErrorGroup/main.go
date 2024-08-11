package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg := errgroup.Group{}

	eg.Go(func() error {
		return getUrl("https://blog.kennycoder.io")
	})

	eg.Go(func() error {
		return getUrl("https://google.com")
	})

	if err := eg.Wait(); err != nil {
		log.Fatalf("got error %v ", err)
	}
	fmt.Println(" No error occured ")
}

func getUrl(url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fail to get page: %s, wrong statusCode: %d", url, resp.StatusCode)
	}

	log.Printf("success get page %s", url)
	return nil
}
