package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	_ "net/http/pprof" //nolint
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"

	"github.com/bend-is/task3/pkg/textprocessor"
)

const (
	minWordLength   = 3
	maxThreads      = 12
	shutdownTimeout = 5
	debug           = "DEBUG"
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
	if os.Getenv("ENV") == debug {
		if err := startCPUProf(); err == nil {
			defer pprof.StopCPUProfile()
		}
	}

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

	serv := &http.Server{Addr: ":8080", Handler: mux}

	go func() {
		<-ctxHnd.Done()
		if os.Getenv("ENV") == debug {
			saveMemProfile()
		}
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
	serv := &http.Server{Addr: ":9000", Handler: nil}

	go func() {
		<-ctx.Done()
		ctxStd, cancel := context.WithTimeout(context.Background(), shutdownTimeout*time.Second)
		defer cancel()

		log.Println("Shutting down the HTTP server on port " + serv.Addr)
		if err := serv.Shutdown(ctxStd); err != nil {
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

func startCPUProf() error {
	cpuProfile, err := os.Create("cpu.prof")
	if err != nil {
		log.Println(err)
		return err
	}

	if err = pprof.StartCPUProfile(cpuProfile); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func saveMemProfile() {
	f, err := os.Create("mem.prof")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	runtime.GC()
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Println(err)
	}
}
