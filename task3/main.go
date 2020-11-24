package main

import (
	"bufio"
	"fmt"
	"github.com/bend-is/task3/pkg/sortedmap"
	"io"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

const filePath = "assets/text.txt"

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error occurred whlie opening file %v: %v", filePath, err)
		return
	}
	defer f.Close()

	sMap, err := CountWords(f)
	if err != nil {
		fmt.Printf("error occurred whlie regexp compling: %v", err)
		return
	}

	fmt.Println(sMap.GetTop10())
}

func CountWords(r io.Reader) (*sortedmap.SortedMap, error) {
	sMap := sortedmap.New()
	bufR := bufio.NewReader(r)
	re, err := regexp.Compile(`[a-z]+'?[a-z]+`)
	if err != nil {
		return nil, err
	}

	for line, err := bufR.ReadString('\n'); err != io.EOF; line, err = bufR.ReadString('\n') {
		line = strings.ToLower(strings.TrimSpace(line))

		for _, sent := range strings.Split(line, ".") {
			if len(sent) == 0 {
				continue
			}
			words := strings.Fields(sent)
			if len(words) < 2 {
				continue
			}
			for _, word := range words[1 : len(words)-1] {
				word = re.FindString(word)
				if utf8.RuneCountInString(strings.ReplaceAll(word, "'", "")) > 3 {
					sMap.IncrementCount(word)
				}
			}
		}
	}

	return sMap, nil
}
