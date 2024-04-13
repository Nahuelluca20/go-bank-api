package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{listenAddr: listenAddr}
}

func (s *APIServer) Run() {

	router := mux.NewRouter()
	router.HandleFunc("/account", s.handleAccount)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) {
	// Start the server
	return nil
}

func (s *APIServer) handleGetAcoount(w http.ResponseWriter, r *http.Request) {
	// Start the server
	return nil
}

func (s *APIServer) handleCreateAcoount(w http.ResponseWriter, r *http.Request) {
	// Start the server
	return nil
}

func (s *APIServer) handleDeleteAcoount(w http.ResponseWriter, r *http.Request) { // Start the server
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) {
	// Start the server
	return nil
}
