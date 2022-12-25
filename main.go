package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"
	method := "GET"
	client := &http.Client{}
	count := 100

	var wg sync.WaitGroup

	for i := 1; i <= count; i++ {
		wg.Add(1)
		url := fmt.Sprintf("%s/%d", url, i)
		go func() {
			defer wg.Done()
			printData(url, method, client)
		}()
	}
	wg.Wait()
}

func printData(url string, method string, client *http.Client) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
