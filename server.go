package main

import (
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type DKeyServer struct {
	dkey *DKey
}

func (dk *DKeyServer) GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, ok := vars["key"]
	if !ok {
		http.Error(w, "no key provided", http.StatusBadRequest)
		return
	}

	value, err := dk.dkey.Get(key)
	if errors.Is(err, ErrorNoSuchKey) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(value))
}

func (dk *DKeyServer) PutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, ok := vars["key"]
	if !ok {
		http.Error(w, "no key provided", http.StatusBadRequest)
		return
	}

	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = dk.dkey.Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (dk *DKeyServer) DeleteHandler(w http.ResponseWriter, r *http.Request) {

}

func (dk *DKeyServer) MakeHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/{key}", dk.GetHandler).Methods("GET")
	r.HandleFunc("/{key}", dk.PutHandler).Methods("PUT")
	r.HandleFunc("/{key}", dk.DeleteHandler).Methods("DELETE")
	return r
}

func NewDKServer() (*DKeyServer, error) {
	d, err := NewDKey()
	if err != nil {
		return nil, err
	}
	return &DKeyServer{d}, nil
}

func main() {
	s, err := NewDKServer()
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8080", s.MakeHandler())
}
