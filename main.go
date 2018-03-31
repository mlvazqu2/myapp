package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var doYouWantToStay bool = true

func main() {
	reader := bufio.NewReader(os.Stdin)
	for doYouWantToStay {
		fmt.Print("Would you like the quote of the day? ")
		userResp, _ := reader.ReadString('\n')
		if strings.Contains("no", strings.ToLower(userResp)) {
			fmt.Println("Bye Felicia")
			doYouWantToStay = false
			continue
		}
	}
}
