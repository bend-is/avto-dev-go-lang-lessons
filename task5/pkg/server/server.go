package server

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof" //nolint
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bend-is/task3/pkg/textprocessor"
)

const (
	minWordLength   = 3
	maxThreads      = 12
	shutdownTimeout = 5
)

type Handler struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	processor  *textprocessor.TextProcessor
}

type TextRequest struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

func (h *Handler) text(w http.ResponseWriter, r *http.Request) {
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

	h.processor.CountWordsFromString(t.Text, t.Number)

	if _, err := w.Write([]byte(`{"result": true}`)); err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) stat(w http.ResponseWriter, r *http.Request) {
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

	stor := h.processor.Storage()
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

func (h *Handler) stop(w http.ResponseWriter, r *http.Request) {
	h.cancelFunc()

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(`{"result": true}`)); err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}

func StartServer(ctx context.Context) error {
	ctxHnd, cancelFunc := context.WithCancel(ctx)

	h := &Handler{
		ctx:        ctxHnd,
		cancelFunc: cancelFunc,
		processor:  textprocessor.New(minWordLength, maxThreads),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/text", h.text)
	mux.HandleFunc("/stat/", h.stat)
	mux.HandleFunc("/stop", h.stop)

	serv := &http.Server{Addr: "localhost:8080", Handler: mux}

	go func() {
		<-ctxHnd.Done()
		ctxStd, cancel := context.WithTimeout(context.Background(), shutdownTimeout*time.Second)
		defer cancel()

		log.Println("Shutting down the HTTP server on port " + serv.Addr)
		if err := serv.Shutdown(ctxStd); err != nil {
			log.Println(err)
		}
	}()

	log.Println("Http server run on http://localhost:8080")
	return serv.ListenAndServe()
}

func StartPprofServer(ctx context.Context) error {
	address := "localhost:9000"
	serv := &http.Server{Addr: address, Handler: nil}

	go func() {
		<-ctx.Done()

		if os.Getenv("ENV") == "DEBUG" {
			saveProfile("http://" + address + "/debug/pprof/profile")
		}

		ctxStd, cancel := context.WithTimeout(context.Background(), shutdownTimeout*time.Second)
		defer cancel()

		log.Println("Shutting down the HTTP server on port " + serv.Addr)
		if err := serv.Shutdown(ctxStd); err != nil {
			log.Println(err)
		}
	}()

	log.Println("Pprof server run on http://" + address)
	return serv.ListenAndServe()
}

func serverError(w http.ResponseWriter, err error) {
	log.Println(err)
	if _, wErr := w.Write([]byte(err.Error())); wErr != nil {
		log.Println(wErr)
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func saveProfile(profileURL string) {
	res, err := http.Get(profileURL) //nolint
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	f, err := os.Create("./profile")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	if _, err := io.Copy(f, res.Body); err != nil {
		log.Println(err)
	}
}
