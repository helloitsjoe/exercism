package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

const MBTA_URL = "https://api-v3.mbta.com"

func fetchMbta(endpoint string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get(MBTA_URL + endpoint)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bodyBytes))
}

func main() {
	println("What MBTA line is your stop on?")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	var wg sync.WaitGroup
	wg.Add(2)
	go fetchMbta("/lines", &wg)
	go fetchMbta("/stops", &wg)
	wg.Wait()

	println("What direction are you going?")
	println("What stop?")
	println(text)
}
