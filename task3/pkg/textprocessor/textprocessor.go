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

var reg = regexp.MustCompile(`[^a-zA-Z]`)

type TextProcessor struct {
	storage    *sortedmap.SortedMap
	wordLength int
	maxThreads int
}

func New(wordLength, maxThreads int) *TextProcessor {
	return &TextProcessor{
		storage:    sortedmap.New(),
		wordLength: wordLength,
		maxThreads: maxThreads,
	}
}

// CountWords Read words from source and add it to SortedMap with it appearance order and count.
func (tp *TextProcessor) CountWordsFromSource(source io.Reader) {
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(source)

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
			tp.CountWordsFromString(line, lineNumber)
		}(line, lineNumber)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	wg.Wait()
}

func (tp *TextProcessor) CountWordsFromString(text string, textOrder int) {
	wordPosition := 0
	text = strings.ReplaceAll(text, "\n", ". ")
	for _, sent := range strings.Split(text, ". ") {
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
				tp.storage.AddStopWord(word)
				continue
			}

			if !tp.storage.IsStopWord(word) {
				tp.storage.IncrementCount(word, sortedmap.NewItemOrder(textOrder, wordPosition))
			}
		}
	}
}

func (tp *TextProcessor) Storage() *sortedmap.SortedMap {
	return tp.storage
}
