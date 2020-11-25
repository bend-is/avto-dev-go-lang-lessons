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
	re := regexp.MustCompile(`[a-z]+'?[a-z]+`)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		words := strings.Fields(strings.ToLower(line))

		for i, dirtyWord := range words {
			// Do not run regexp if dirtyWord already to short even with symbols.
			if utf8.RuneCountInString(dirtyWord) <= tp.wordLength {
				continue
			}
			cleanWord := re.FindString(dirtyWord)
			// Replace for words like it's cause we compare only letters lengths and symbol ' must be excluded.
			if utf8.RuneCountInString(strings.Replace(cleanWord, "'", "", 1)) <= tp.wordLength {
				continue
			}
			// Add first or last word cause it definitely a start or end of the sentence.
			if i == 0 || i == len(words)-1 {
				tp.stopWords[cleanWord] = true
				continue
			}
			// Add end of the sentence.
			if strings.Contains(dirtyWord, ".") {
				tp.stopWords[cleanWord] = true
				// Add possible next word that must be a start of sentence.
				if len(words) > i+1 {
					tp.stopWords[re.FindString(words[i+1])] = true
				}
				continue
			}
			if tp.stopWords[cleanWord] {
				sMap.Delete(cleanWord)
				continue
			}

			sMap.IncrementCount(dirtyWord)
		}
	}

	return sMap
}
