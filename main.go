package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// quoteResponse := QuoteResponse{}
	response, err := http.Get("http://quotes.rest/qod.json")
	if err != nil {
		fmt.Printf("Failed to get yo quote, this why: %s", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	var quotes QuoteResponse
	json.Unmarshal(data, &quotes)

	if quotes.Success == nil {
		fmt.Printf("%s\n", quotes.Contents.Quotes[0].Quote)
		fmt.Printf("    - %s\n", quotes.Contents.Quotes[0].Author)
	} else {
		var err SomeError
		json.Unmarshal(data, &err)
		fmt.Printf("Status Code: %d\n%s\n", err.Error.Code, err.Error.Message)
	}
}
