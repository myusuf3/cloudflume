package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sites struct {
}

func getDomains() (*Sites, error) {

	return nil, nil
}

func getExternal() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content)
}

func main() {
	// domains, err := getDomains()

	// get ip
	ip := getExternal()
	fmt.Println("external ip: ", ip)

	// for domains
}
