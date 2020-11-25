package textprocessor

import (
	"io"
	"os"
	"strings"
	"testing"
)

type TestCase struct {
	text           io.Reader
	expectedWords  []string
	expectedCounts []int
}

func TestCountWords(t *testing.T) {
	f, _ := os.Open("testdata/text.txt")
	testCases := []TestCase{
		{
			strings.NewReader(""),
			[]string{"", "", "", "", "", "", "", "", "", ""},
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			strings.NewReader(`
				Some words in large string. Maybe it's not so large as you think
				Some words in large string. Maybe not so large as you think.
				Some words in large string. Maybe not so large as you think
			`),
			[]string{"words", "large", "", "", "", "", "", "", "", ""},
			[]int{3, 6, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			strings.NewReader(`
				Hello hello again hello my friend. Blast blast again blast blast. Hell hell hell hell yeah!
				Goodbye blast Goodbye Goodbye my friend. Shine Shine Shine. Home home home home!
				She'll She'll - She'll: She'll. Blast blast again blast blast. Chair blast Chair Chair hell yeah!
				Repeat repeat repeat repeat repeat repeat again and repeat repeat!!
				And this word won't be on our list. But list list will will be
			`),
			[]string{"again", "this", "word", "won't", "will", "", "", "", "", ""},
			[]int{4, 1, 1, 1, 2, 0, 0, 0, 0, 0},
		},
		{
			f,
			[]string{"like", "told", "looked", "went", "love", "want", "into", "took", "heard", "can't"},
			[]int{34, 29, 38, 61, 28, 28, 29, 20, 16, 17},
		},
	}

	for _, tt := range testCases {
		res := New(tt.text).CountWords()
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
