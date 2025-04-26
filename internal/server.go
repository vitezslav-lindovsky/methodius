package methodius

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const CYRIL = `
        ____  
      .-" + "-.
     /   .--.  \    
     |  | () |  |   <- Cyril thinking deeply
     |  |____|  |
     \   '--'  _/
     /'-.____.-'\   
   / /   ||    \ \  
  /_/    ||     \_\  
   ||____||____||  
  /_/    ||     \_\  
 |__\____||____/__|    <- Humble robes, monk-mode
   /    /||\    \     
  /___.' || '.___\
    ||   ||   ||
    []   []   []

"Hmm... perhaps *ъ* should be softer."

	Academic and ascetic
	Resting  but  divine

https://www.youtube.com/watch?v=Nl6bWvzLaHs

`

const (
	Year         = 444
	ActionRead   = "read"
	ActionCreate = "create"
	ActionUpdate = "update"
	ActionRemove = "remove"
	ActionQuit   = "quit"
	MethodQuit   = "QUIT"
)

type Server struct {
	store   *KeyValueStore
	verbose bool
	actions map[string]string
}

func NewServer(store *KeyValueStore, verbose bool, methodsToActions map[string]string) *Server {
	return &Server{
		store:   store,
		verbose: verbose,
		actions: methodsToActions,
	}
}

func (s *Server) HandleRequest(w http.ResponseWriter, r *http.Request) {
	action := s.actions[r.Method]

	switch action {
	case ActionRead:
		s.handleRead(w, r)
	case ActionCreate:
		s.handleCreate(w, r)
	case ActionUpdate:
		s.handleUpdate(w, r)
	case ActionRemove:
		s.handleRemove(w, r)
	case ActionQuit:
		s.handleQuit(w)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleQuit(w http.ResponseWriter) {
	s.logVerbose("Quitting")

	http.Error(w, CYRIL, http.StatusOK)

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}

	// Relax, this code is just for fun.
	time.Sleep(3 * time.Second)
	os.Exit(Year)
}

func (s *Server) handleRead(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]

	if key == "" {
		s.handleList(w, r)
		return
	}

	s.logVerbose("Doing Read for %s", r.Method)
	value, exists := s.store.Get(key)

	if !exists {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	item := map[string]string{key: value}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func (s *Server) handleList(w http.ResponseWriter, r *http.Request) {
	s.logVerbose("Doing List for %s", r.Method)
	allItems := s.store.GetAll()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allItems)
}

func (s *Server) handleCreate(w http.ResponseWriter, r *http.Request) {
	s.logVerbose("Doing Create for %s", r.Method)

	key := r.URL.Path[1:]

	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	value, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = s.store.Set(key, string(value))
	if err != nil {
		if err.Error() == "key already exists" {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Created [%s:%s]", key, value)
}

func (s *Server) handleUpdate(w http.ResponseWriter, r *http.Request) {
	s.logVerbose("Doing Update for %s", r.Method)

	key := r.URL.Path[1:]

	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	value, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = s.store.Update(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Updated key: %s", key)
}

func (s *Server) handleRemove(w http.ResponseWriter, r *http.Request) {
	s.logVerbose("Doing Remove for %s", r.Method)

	key := r.URL.Path[1:]

	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	err := s.store.Delete(key)
	if err != nil {
		if err.Error() == "key not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted key: %s", key)
}

func (s *Server) logVerbose(msg string, a ...any) {
	if s.verbose {
		fmt.Printf(msg+"\n", a...)
	}
}
