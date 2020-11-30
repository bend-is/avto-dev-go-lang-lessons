package textprocessor

import (
	"bufio"
	"io"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/bend-is/task3/pkg/sortedmap"
)

type TextProcessor struct {
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
	scanner := bufio.NewScanner(tp.source)
	reg := regexp.MustCompile(`[^a-zA-Z]`)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		for _, sent := range strings.Split(line, ". ") {
			words := strings.Fields(strings.ToLower(sent))

			for i, word := range words {
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
					tp.stopWords[word] = true
					continue
				}
				// If exist in stop words - remove it.
				if tp.stopWords[word] {
					sMap.Delete(word)
					continue
				}

				sMap.IncrementCount(word)
			}
		}
	}

	return sMap
}
