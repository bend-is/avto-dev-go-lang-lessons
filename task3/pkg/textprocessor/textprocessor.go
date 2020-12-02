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

const (
	// Need to keep order. Thinking that can't be more then 1000 words in line.
	orderMultiplier = 1000
	// Number of max goroutines.
	maxGoroutines = 12
)

type TextProcessor struct {
	sync.Mutex
	source     io.Reader
	wordLength int
}

func New(reader io.Reader, wordLength int) *TextProcessor {
	return &TextProcessor{
		source:     reader,
		wordLength: wordLength,
	}
}

// CountWords Read words from source and add it to SortedMap with it appearance order and count.
func (tp *TextProcessor) CountWords() *sortedmap.SortedMap {
	var wg sync.WaitGroup
	sMap := sortedmap.New()
	scanner := bufio.NewScanner(tp.source)
	reg := regexp.MustCompile(`[^a-zA-Z]`)

	wordCh, stopCh := tp.listenStopWords(sMap)
	guardCh := make(chan struct{}, maxGoroutines)

	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}
		wg.Add(1)
		guardCh <- struct{}{}
		go func(lineNumber int) {
			defer wg.Done()
			defer func() { <-guardCh }()

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
						wordCh <- word
						continue
					}

					tp.Lock()
					sMap.IncrementCount(word, lineNumber*orderMultiplier+wordPosition)
					tp.Unlock()
				}
			}
		}(lineNumber)
	}
	wg.Wait()
	// Sync all stopWords and close stopWord listener.
	stopCh <- struct{}{}

	return sMap
}

// nolint
func (tp *TextProcessor) listenStopWords(sMap *sortedmap.SortedMap) (chan<- string, chan<- struct{}) {
	wordCh := make(chan string)
	stopCh := make(chan struct{})

	go func() {
		for {
			select {
			case word := <-wordCh:
				sMap.AddStopWord(word)
			case <-stopCh:
				break
			}
		}
	}()

	return wordCh, stopCh
}
