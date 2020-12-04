package textprocessor

import (
	"bufio"
	"io"
	"regexp"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/bend-is/task3/pkg/sortedmap"
)

type TextProcessor struct {
	source     io.Reader
	wordLength int
	maxThreads int
}

func New(reader io.Reader, wordLength, maxThreads int) *TextProcessor {
	return &TextProcessor{
		source:     reader,
		wordLength: wordLength,
		maxThreads: maxThreads,
	}
}

// CountWords Read words from source and add it to SortedMap with it appearance order and count.
func (tp *TextProcessor) CountWords() *sortedmap.SortedMap {
	var wg sync.WaitGroup
	sMap := sortedmap.New()
	scanner := bufio.NewScanner(tp.source)
	reg := regexp.MustCompile(`[^a-zA-Z]`)

	guardCh := make(chan struct{}, tp.maxThreads)
	defer close(guardCh)

	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}
		wg.Add(1)
		guardCh <- struct{}{}
		go func(line string, lineNumber int) {
			defer func() { wg.Done(); <-guardCh }()

			wordPosition := 0
			for _, sent := range strings.Split(line, ". ") {
				words := strings.Fields(strings.ToLower(sent))

				for i, word := range words {
					wordPosition++
					// Do not run regexp if dirtyWord already to short even with symbols.
					if utf8.RuneCountInString(word) <= tp.wordLength {
						continue
					}
					word = reg.ReplaceAllString(word, "")
					if utf8.RuneCountInString(word) <= tp.wordLength {
						continue
					}
					// Add first or last word cause it definitely a start or end of the sentence.
					if i == 0 || i == len(words)-1 {
						sMap.AddStopWord(word)
						continue
					}

					if !sMap.IsStopWord(word) {
						sMap.IncrementCount(word, sortedmap.NewItemOrder(lineNumber, wordPosition))
					}
				}
			}
		}(line, lineNumber)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	wg.Wait()

	return sMap
}
