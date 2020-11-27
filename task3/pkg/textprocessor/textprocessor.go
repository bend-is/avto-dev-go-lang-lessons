package textprocessor

import (
	"bufio"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/bend-is/task3/pkg/sortedmap"
)

type TextProcessor struct {
	source     io.Reader
	stopWords  map[string]bool
	wordLength int
}

func New(reader io.Reader) *TextProcessor {
	return &TextProcessor{
		source:     reader,
		stopWords:  make(map[string]bool),
		wordLength: 3,
	}
}

// WordLength sets the word length for CountWords function
func (tp *TextProcessor) WordLength(length int) {
	tp.wordLength = length
}

// CountWords Read words from source and add it to SortedMap with it appearance order and count.
func (tp *TextProcessor) CountWords() *sortedmap.SortedMap {
	sMap := sortedmap.New()
	scanner := bufio.NewScanner(tp.source)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		words := strings.Fields(strings.ToLower(line))

		for i, word := range words {
			// Do not run regexp if dirtyWord already to short even with symbols.
			if utf8.RuneCountInString(word) <= tp.wordLength {
				continue
			}
			word = strings.TrimFunc(word, func(r rune) bool {
				return !unicode.IsLetter(r)
			})
			// For case like she'll.
			word = strings.ReplaceAll(word, "'", "")
			if utf8.RuneCountInString(word) <= tp.wordLength {
				continue
			}
			// Add first or last word cause it definitely a start or end of the sentence.
			if i == 0 || i == len(words)-1 {
				tp.stopWords[word] = true
				continue
			}
			// Add end of the sentence.
			if strings.Contains(words[i], ".") {
				tp.stopWords[word] = true
				continue
			}
			// Add sentence first word.
			if strings.Contains(words[i-1], ".") {
				tp.stopWords[word] = true
			}
			// If exist in stop words - remove it.
			if tp.stopWords[word] {
				sMap.Delete(word)
				continue
			}

			sMap.IncrementCount(word)
		}
	}

	return sMap
}
