package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCountWords(t *testing.T) {
	string := `
		Some words in large string. Maybe not so large as you think.
		Some words in large string. Maybe not so large as you think.
		Some words in large string. Maybe not so large as you think.
	`

	res, err := CountWords(strings.NewReader(string))
	if err != nil {
		t.Errorf("%v", err)
	}

	fmt.Println(res.GetTop10())
}
