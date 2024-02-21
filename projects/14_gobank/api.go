package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store Storage
}

func NewAPIServer(listenAdd string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAdd,
		store: store,
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/login", makeHTTPHandleFunc((s.handleLogin)))
	router.HandleFunc("/accounts", makeHTTPHandleFunc(s.handleAccounts))
	router.HandleFunc("/accounts/{id}", WithJWTAuth(makeHTTPHandleFunc(s.handleAccount), s.store))
	router.HandleFunc("/transfer", makeHTTPHandleFunc(s.handleTransfer))

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	
	return WriteJSON(w, http.StatusOK, req)
}

func (s *APIServer) handleAccounts(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleFetchAccounts(w, r)
	} else if r.Method == "POST" {
			return s.handleCreateAccount(w, r)
	}
			
	return fmt.Errorf("method not allowed %s", r.Method)
}
		
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	} else if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	} else if r.Method == "PATCH" {
		return s.handleUpdateAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleUpdateAccount(w http.ResponseWriter, r *http.Request) error {
	updateAccountReq := &UpdateAccountRequest{}
	id, err := getID(r)

	if err != nil {
		return err
	}

	if err := json.NewDecoder(r.Body).Decode(&updateAccountReq); err != nil {
		return err
	}

	defer r.Body.Close()

	account := &Account{
		FirstName: updateAccountReq.FirstName,
		LastName:  updateAccountReq.LastName,
	}

	if err := s.store.UpdateAccount(id, account); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, account)
}


func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)

	if err != nil {
		return err
	}

	accounts, err := s.store.GetAccountById(id)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountReq := &CreateAccountRequest{} 

	if err := json.NewDecoder(r.Body).Decode(&createAccountReq); err != nil {
		return err
	}

	defer r.Body.Close()

	account, err := NewAccount(createAccountReq.FirstName, createAccountReq.LastName, createAccountReq.Password)

	if err != nil {
		return err
	}

	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)

	if err != nil {
		return err
	}

	s.store.DeleteAccount(id)

	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	transferReq := &TransferRequest{}

	if err := json.NewDecoder(r.Body).Decode(transferReq); err != nil {
		return err
	}

	defer r.Body.Close()



	return WriteJSON(w, http.StatusOK, transferReq)
}

func (s *APIServer) handleFetchAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.FetchAccounts()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, accounts)
}

func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)

	return id, err
}



