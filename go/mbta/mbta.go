package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

const MBTA_URL = "https://api-v3.mbta.com"

func main() {
	println("What MBTA line is your stop on?")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	res, err := http.Get(MBTA_URL + "/lines")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bodyBytes))

	println("What direction are you going?")
	println("What stop?")
	println(text)
}
