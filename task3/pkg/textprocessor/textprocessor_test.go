package textprocessor

import (
	"os"
	"strings"
	"testing"
)

type TestCase struct {
	text           string
	wordsLength    int
	expectedWords  []string
	expectedCounts []int
}

func TestCountWords(t *testing.T) {
	testCases := []TestCase{
		{
			"",
			3,
			[]string{"", "", "", "", "", "", "", "", "", ""},
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			`
				Some words! in large string. Maybe it's not so large!! as you think
				Some words in large string. Maybe not so -large as you think.
				Some words? in large string. Maybe not so ?large as you think?
			`,
			3,
			[]string{"words", "large", "", "", "", "", "", "", "", ""},
			[]int{3, 6, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			`
				Hello hello again! Gree244tings my friend. Hello again I told you that.
				Words. that. not. be. in. this. list. cheers cheers! Happy hell yeah!
				Repeat repeat inlistedword repeat repeat repeat repeat again and repeat repeat!!
				And this 'repeat' word won't be on our list. But golang golang will will be in
			`,
			3,
			[]string{"again", "greetings", "told", "happy", "hell", "inlistedword", "word", "wont", "golang", "will"},
			[]int{3, 1, 1, 1, 1, 1, 1, 1, 2, 2},
		},
		{
			`
				Hello hello again! Greetings my friend. Hello again I told you that.
				Words. that. not. be. in. this. list. cheers cheers! Happy hell yeah!
				Repeat repeat inlistedword repeat repeat repeat repeat again and repeat repeat!!
				And this 'repeat' word won't be on our list. But golang golang will will be in
			`,
			5,
			[]string{"greetings", "inlistedword", "golang", "", "", "", "", "", "", ""},
			[]int{1, 1, 2, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tt := range testCases {
		proc := New(tt.wordsLength, 12)
		proc.CountWordsFromSource(strings.NewReader(tt.text))
		for i, v := range proc.Storage().GetTop(10) {
			if tt.expectedWords[i] != v {
				t.Fatalf("unexpected word at index %d: wont %s - got %s", i, tt.expectedWords[i], v)
			}
			if tt.expectedCounts[i] != proc.Storage().GetCount(v) {
				t.Fatalf("unexpected count for '%s': wont %d - got %d", v, tt.expectedCounts[i], proc.Storage().GetCount(v))
			}
		}
	}
}

func TestCountWordsOnFile(t *testing.T) {
	expectedWords := []string{"like", "told", "looked", "marry", "went", "love", "want", "into", "took", "cant"}
	expectedCount := []int{35, 29, 39, 21, 62, 32, 29, 29, 20, 18}
	f, _ := os.Open("testdata/text.txt")
	defer f.Close()

	proc := New(3, 12)
	proc.CountWordsFromSource(f)
	for i, v := range proc.Storage().GetTop(10) {
		if expectedWords[i] != v {
			t.Fatalf("unexpected word at index %d: wont %s - got %s", i, expectedWords[i], v)
		}
		if expectedCount[i] != proc.Storage().GetCount(v) {
			t.Fatalf("unexpected count for '%s': wont %d - got %d", v, expectedCount[i], proc.Storage().GetCount(v))
		}
	}
}

func BenchmarkCountWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		f, _ := os.Open("testdata/text.txt")
		b.StartTimer()
		proc := New(3, 12)
		proc.CountWordsFromSource(f)
	}
}
