package main

import (
	"fmt"
	"github.com/mark2b/faceplusplus-client"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		panic("Wrong arguments count. Execute with API_KEY, API_SECRET")
	}
	client := faceplusplus_client.NewClient(args[1], args[2])
	if imageData, err := ioutil.ReadFile("examples/demo-pic.jpg"); err == nil {
		if response, err := client.DetectHumanBodySegmentWithImageData(imageData); err == nil {
			fmt.Printf("response: %v", response)
		} else {
			println("error", err.Error())
		}
	}

}
