package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Would you like the quote of the day? ")
	userResp, _ := reader.ReadString('\n')
	userResp = userResp[:len(userResp)-1]
	if strings.Contains("no", strings.ToLower(userResp)) || userResp == "no" {
		fmt.Println("Well... have a good one. :)")
	} else {
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

}
