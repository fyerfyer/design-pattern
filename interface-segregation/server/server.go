package server

import (
	"fmt"
	"interface/updater"
	"net/http"
	"sync"
)

type SimpleHttpServer struct {
	host    string
	port    int
	viewers map[string][]updater.Viewer
	mu      sync.RWMutex
}

func NewSimpleHttpServer(host string, port int) *SimpleHttpServer {
	return &SimpleHttpServer{
		host:    host,
		port:    port,
		viewers: make(map[string][]updater.Viewer),
	}
}

func (s *SimpleHttpServer) AddViewer(urlDirectory string, viewer updater.Viewer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.viewers[urlDirectory] = append(s.viewers[urlDirectory], viewer)
}

func (s *SimpleHttpServer) Run() {
	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		s.mu.RLock()
		defer s.mu.RUnlock()

		viewers, exists := s.viewers["/config"]
		if !exists {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		for _, viewer := range viewers {
			fmt.Fprintln(w, viewer.OutputInPlainText())
		}
	})

	http.ListenAndServe(fmt.Sprintf("%s:%d", s.host, s.port), nil)
}
