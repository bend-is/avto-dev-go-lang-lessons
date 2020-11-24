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

	fmt.Printf("Most repetead words:\n")
	for _, v := range sMap.GetTop10() {
		fmt.Printf("'%v' was repeated: %v time\n", v, sMap.GetCount(v))
	}
}

//CountWords Read words from reader and add it to SortedMap with it appearance order and count.
func CountWords(r io.Reader) (*sortedmap.SortedMap, error) {
	sMap := sortedmap.New()
	bufR := bufio.NewReader(r)
	re, err := regexp.Compile(`[a-z]+'?[a-z]+`)
	if err != nil {
		return nil, err
	}

	for line, err := bufR.ReadString('\n'); err != io.EOF; line, err = bufR.ReadString('\n') {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		words := strings.Fields(strings.ToLower(line))

		for i := 1; i < len(words)-1; i++ {
			word := words[i]
			if strings.Contains(word, ".") {
				i++ // Skip this word and the next cause it will be the start of sentence.
				continue
			}
			// Do not run regexp if word already to short even with symbols.
			if utf8.RuneCountInString(word) < 4 {
				continue
			}
			// Clean word.
			word = re.FindString(word)
			// Replace - for words like it's. Cause we need more then 3 letters lengths and symbols must be excluded.
			if utf8.RuneCountInString(strings.Replace(word, "'", "", 1)) > 3 {
				sMap.IncrementCount(word)
			}
		}
	}

	return sMap, nil
}
