package textprocessor

import (
	"io"
	"os"
	"strings"
	"testing"
)

type TestCase struct {
	text           io.Reader
	wordsLength    int
	expectedWords  []string
	expectedCounts []int
}

func TestCountWords(t *testing.T) {
	f, _ := os.Open("testdata/text.txt")
	defer f.Close()
	testCases := []TestCase{
		{
			strings.NewReader(""),
			3,
			[]string{"", "", "", "", "", "", "", "", "", ""},
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			strings.NewReader(`
				Some words! in large string. Maybe it's not so large!! as you think
				Some words in large string. Maybe not so -large as you think.
				Some words? in large string. Maybe not so ?large as you think?
			`),
			3,
			[]string{"words", "large", "", "", "", "", "", "", "", ""},
			[]int{3, 6, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			strings.NewReader(`
				Hello hello again! Greetings my friend. Hello again I told you that.
				Words. that. not. be. in. this. list. cheers cheers! Happy hell yeah!
				Repeat repeat inlistedword repeat repeat repeat repeat again and repeat repeat!!
				And this 'repeat' word won't be on our list. But golang golang will will be in
			`),
			3,
			[]string{"again", "greetings", "told", "happy", "hell", "inlistedword", "word", "won't", "golang", "will"},
			[]int{3, 1, 1, 1, 1, 1, 1, 1, 2, 2},
		},
		{
			strings.NewReader(`
				Hello hello again! Greetings my friend. Hello again I told you that.
				Words. that. not. be. in. this. list. cheers cheers! Happy hell yeah!
				Repeat repeat inlistedword repeat repeat repeat repeat again and repeat repeat!!
				And this 'repeat' word won't be on our list. But golang golang will will be in
			`),
			5,
			[]string{"greetings", "inlistedword", "golang", "", "", "", "", "", "", ""},
			[]int{1, 1, 2, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			f,
			3,
			[]string{"looked", "marry", "went", "love", "want", "into", "took", "heard", "something", "can't"},
			[]int{39, 21, 62, 32, 29, 29, 20, 16, 16, 18},
		},
	}

	for _, tt := range testCases {
		tp := New(tt.text)
		tp.WordLength(tt.wordsLength)
		res := tp.CountWords()
		for i, v := range res.GetTop(10) {
			if tt.expectedWords[i] != v {
				t.Fatalf("unexpected word at index %d: wont %s - got %s", i, tt.expectedWords[i], v)
			}
			if tt.expectedCounts[i] != res.GetCount(v) {
				t.Fatalf("unexpected count for '%s': wont %d - got %d", v, tt.expectedCounts[i], res.GetCount(v))
			}
		}
	}
}

func BenchmarkCountWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		f, _ := os.Open("testdata/text.txt")
		b.StartTimer()
		_ = New(f).CountWords()
	}
}
