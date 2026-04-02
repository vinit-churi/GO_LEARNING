//go:build !solution
// +build !solution

package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type StatusService struct {
	mu     sync.Mutex
	values map[string]string
}

func NewStatusService() *StatusService {
	return &StatusService{values: map[string]string{}}
}

func (s *StatusService) SetStatus(service, status string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[service] = status
}

func (s *StatusService) GetStatus(service string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	value, ok := s.values[service]
	return value, ok
}

func (s *StatusService) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service := r.URL.Query().Get("service")
		status, _ := s.GetStatus(service)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"service": service,
			"status":  status,
		})
	}
}

func main() {}
