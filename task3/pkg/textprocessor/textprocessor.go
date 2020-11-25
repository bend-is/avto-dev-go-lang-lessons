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
	wordLength int
}

func New(reader io.Reader) *TextProcessor {
	return &TextProcessor{
		source:     reader,
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

		for i := 1; i < len(words)-1; i++ {
			word := words[i]
			if strings.Contains(word, ".") {
				i++ // Skip this word and the next cause it will be the start of sentence.
				continue
			}
			// Do not run regexp if word already to short even with symbols.
			if utf8.RuneCountInString(word) <= tp.wordLength {
				continue
			}
			// Clean word.
			word = re.FindString(word)
			// Replace - for words like it's. Cause we need more then 3 letters lengths and symbols must be excluded.
			if utf8.RuneCountInString(strings.Replace(word, "'", "", 1)) > tp.wordLength {
				sMap.IncrementCount(word)
			}
		}
	}

	return sMap
}
