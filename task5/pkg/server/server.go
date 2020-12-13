package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	//nolint
	_ "net/http/pprof"
	"strconv"
	"strings"

	"github.com/bend-is/task3/pkg/textprocessor"
)

const (
	minWordLength = 3
	maxThreads    = 12
)

var (
	processor  *textprocessor.TextProcessor
	cancelFunc context.CancelFunc
)

type TextRequest struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

func text(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	t := TextRequest{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		serverError(w, err)
		return
	}
	if t.Number == 0 || t.Text == "" {
		if _, err := w.Write([]byte("Invalid Request")); err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	processor.CountWordsFromString(t.Text, t.Number)

	if _, err := w.Write([]byte(`{"result": true}`)); err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}

func stat(w http.ResponseWriter, r *http.Request) {
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
		serverError(w, err)
		return
	}
	if _, err = w.Write(jsn); err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}

func stop(w http.ResponseWriter, r *http.Request) {
	cancelFunc()

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(`{"result": true}`)); err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}

func StartServer(ctx context.Context) error {
	ctx, cancelFunc = context.WithCancel(ctx)

	processor = textprocessor.New(minWordLength, maxThreads)

	mux := http.NewServeMux()
	mux.HandleFunc("/text", text)
	mux.HandleFunc("/stat/", stat)
	mux.HandleFunc("/stop", stop)

	serv := &http.Server{Addr: ":8080", Handler: mux}

	go func() {
		<-ctx.Done()
		log.Println("Shutting down the HTTP server...")
		if err := serv.Shutdown(ctx); err != nil {
			log.Println(err)
		}
	}()

	log.Println("Http server run on http://localhost:8080")
	return serv.ListenAndServe()
}

func StartPprofServer(ctx context.Context) error {
	serv := &http.Server{Addr: ":9000", Handler: nil}

	go func() {
		<-ctx.Done()
		log.Println("Shutting down the HTTP Pprof server...")
		if err := serv.Shutdown(ctx); err != nil {
			log.Println(err)
		}
	}()

	log.Println("Pprof server run on http://localhost:9000")
	return serv.ListenAndServe()
}

func serverError(w http.ResponseWriter, err error) {
	log.Println(err)
	if _, wErr := w.Write([]byte(err.Error())); wErr != nil {
		log.Println(wErr)
	}
	w.WriteHeader(http.StatusInternalServerError)
}
