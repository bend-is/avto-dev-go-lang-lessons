package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bend-is/task3/pkg/textprocessor"
)

func main() {
	var topCount, wordLength, maxThreads int
	var filePath string

	flag.IntVar(&topCount, "c", 10, "count of top repeated words")
	flag.IntVar(&wordLength, "wl", 3, "word length less than which words are skipped")
	flag.IntVar(&maxThreads, "t", 12, "max number of threads for text parsing process")
	flag.StringVar(&filePath, "f", "assets/text.txt", "file path for text processing")
	flag.Parse()

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error occurred while opening filePath %s: %s", filePath, err)
		return
	}
	defer f.Close()

	processr := textprocessor.New(wordLength, maxThreads)
	processr.CountWordsFromSource(f)

	fmt.Printf("Most repetead words:\n")
	for _, v := range processr.Storage().GetTop(topCount) {
		if v == "" {
			continue
		}
		fmt.Printf("'%s' was repeated: %d time\n", v, processr.Storage().GetCount(v))
	}
}
