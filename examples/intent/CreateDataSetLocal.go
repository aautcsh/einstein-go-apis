package main

import (
	"os"
	"fmt"
	"github.com/rajdeepd/einstein-go-apis/lang"
)

func main() {
	accessToken := os.Getenv("ACCESSKEY")
	path := "/home/ubuntu/work/einstein-scala-lib/data/CallRouting.csv"
	response, err := lang.CreateIntentDataSet(path, accessToken)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Response: " + string(response))
}