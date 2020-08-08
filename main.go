package main

import (
	"github.com/aws/aws-sdk-go"
	"github.com/aws/aws-sdk-go/service/athena"
)

func main() {
	var strPtrs []*string
	var strs []string = []string{"Go", "Gophers", "Go"}

	// Convert []string to []*string
	strPtrs = aws.StringSlice(strs)

	// Convert []*string to []string
	strs = aws.StringValueSlice(strPtrs)
}
