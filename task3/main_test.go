package main

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
	f, _ := os.Open(filePath)
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
				Hello hello hello my friend. Blast blast blast blast. Hell hell hell hell yeah!
				Goodbye blast Goodbye Goodbye my friend. Shine Shine Shine. Home home home home!
				She'll She'll - She'll: She'll. Blast blast blast blast. Chair blast Chair Chair hell yeah!
				Repeat repeat repeat repeat repeat repeat again and repeat repeat!!
				And this word won't be on our list. But list list will be
			`),
			[]string{"hello", "blast", "hell", "goodbye", "home", "she'll", "chair", "repeat", "again", "list"},
			[]int{2, 6, 4, 2, 2, 2, 2, 6, 1, 2},
		},
		{
			f,
			[]string{"that", "from", "with", "scarlett", "said", "ashley", "melanie", "went", "came", "thought"},
			[]int{126, 68, 128, 146, 200, 49, 63, 62, 43, 46},
		},
	}

	for _, tt := range testCases {
		res, err := CountWords(tt.text)
		if err != nil {
			t.Errorf("%v", err)
		}

		for i, v := range res.GetTop10() {
			if tt.expectedWords[i] != v {
				t.Fatalf("unexpected word at index %v: wont %v - got %v", i, tt.expectedWords[i], v)
			}
			if tt.expectedCounts[i] != res.GetCount(v) {
				t.Fatalf("unexpected count for '%v': wont %v - got %v", v, tt.expectedCounts[i], res.GetCount(v))
			}
		}
	}
}

func BenchmarkCountWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		f, _ := os.Open(filePath)
		b.StartTimer()
		_, _ = CountWords(f)
	}
}
