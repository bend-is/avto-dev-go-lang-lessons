package main

import (
	"fmt"
	"os"

	"github.com/bend-is/task3/pkg/textprocessor"
)

const (
	defaultFilePath = "assets/text.txt"
	defaultTopCount = 10
)

func main() {
	f, err := os.Open(defaultFilePath)
	if err != nil {
		fmt.Printf("error occurred while opening file %s: %s", defaultFilePath, err)
		return
	}
	defer f.Close()

	tp := textprocessor.New(f)
	sMap := tp.CountWords()

	fmt.Printf("Most repetead words:\n")
	for _, v := range sMap.GetTop(defaultTopCount) {
		fmt.Printf("'%s' was repeated: %d time\n", v, sMap.GetCount(v))
	}
}
