package main

import (
	"fmt"
	"github.com/mark2b/faceplusplus-client"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		panic("Wrong arguments count. Execute with API_KEY, API_SECRET")
	}
	client := faceplusplus_client.NewClient(args[1], args[2])
	if response, err := client.DetectBody("examples/demo-pic.jpg", []string{"gender"}); err == nil {
		fmt.Printf("response: %v", response)
	} else {
		println("error", err.Error())
	}
}
