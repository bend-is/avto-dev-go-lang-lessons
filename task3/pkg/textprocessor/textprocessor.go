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
	sync.Mutex
	source     io.Reader
	stopWords  map[string]bool
	wordLength int
}

func New(reader io.Reader, wordLength int) *TextProcessor {
	return &TextProcessor{
		source:     reader,
		stopWords:  make(map[string]bool),
		wordLength: wordLength,
	}
}

// CountWords Read words from source and add it to SortedMap with it appearance order and count.
func (tp *TextProcessor) CountWords() *sortedmap.SortedMap {
	sMap := sortedmap.New()
	wg := new(sync.WaitGroup)
	scanner := bufio.NewScanner(tp.source)
	reg := regexp.MustCompile(`[^a-zA-Z]`)

	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}
		wg.Add(1)
		go func(lineNumber int) {
			defer wg.Done()
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
						tp.addStopWord(word)
						sMap.Delete(word)
						continue
					}
					// If exist in stop words - remove it.
					if tp.existInStopWords(word) {
						sMap.Delete(word)
						continue
					}

					// Magic 100 to keep order.
					sMap.IncrementCount(word, lineNumber*100+wordPosition)
				}
			}
		}(lineNumber)
	}
	wg.Wait()

	return sMap
}

func (tp *TextProcessor) addStopWord(word string) {
	tp.Lock()
	defer tp.Unlock()

	tp.stopWords[word] = true
}

func (tp *TextProcessor) existInStopWords(word string) bool {
	tp.Lock()
	defer tp.Unlock()

	return tp.stopWords[word]
}
