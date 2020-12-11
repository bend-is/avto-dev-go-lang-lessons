package main

import (
	"encoding/json"
	"github.com/prometheus/common/log"
	"net/http"
	"strconv"
	"strings"

	"github.com/bend-is/task3/pkg/textprocessor"
)

var processor *textprocessor.TextProcessor

type TextRequest struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

func Text(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	t := TextRequest{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if t.Number == 0 || t.Text == "" {
		_, _ = w.Write([]byte("Invalid Request"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	processor.CountWordsFromString(t.Text, t.Number)
	w.WriteHeader(http.StatusOK)

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"result": true}`))
	if err != nil {
		log.Error(err)
	}
}

func Stat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	strN := strings.TrimPrefix(r.URL.Path, "/stat/")
	n, err := strconv.Atoi(strN)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	stor := processor.Storage()
	res := make(map[string]int, n)

	for _, v := range stor.GetTop(n) {
		if v == "" {
			continue
		}
		res[v] = stor.GetCount(v)
	}

	jsn, err := json.Marshal(&res)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsn)
	if err != nil {
		log.Error(err)
	}
}

func Stop(w http.ResponseWriter, r *http.Request) {
	//
}

func main() {
	processor = textprocessor.New(3, 12)

	mux := http.NewServeMux()

	mux.HandleFunc("/text", Text)
	mux.HandleFunc("/stat/", Stat)
	mux.HandleFunc("/stop", Stop)

	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		panic(err)
	}
}
