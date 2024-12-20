package httpserver

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	//"github.com/gorilla/mux"
)

const defPort = "8000"

type HttpServer struct {
	port    string
	storage StorageItf
}

func Init(storage StorageItf) *HttpServer {
	httpServer := &HttpServer{
		storage: storage,
		port:    defPort,
	}

	if os.Getenv("PORT") != "" {
		httpServer.port = os.Getenv("PORT")
	}
	http.HandleFunc("/api/login", httpServer.LoginAttempt)
	http.HandleFunc("/api/user", httpServer.ReturnLogin)
	http.HandleFunc("/api/logout", httpServer.Logout)
	http.HandleFunc("/api/register", httpServer.Register)
	return httpServer
}

func InitAndStart(storage StorageItf) *HttpServer {
	httpServer := Init(storage)
	httpServer.Start()
	return httpServer
}

func (httpServer *HttpServer) Start() {
	log.Printf("Start HTTP Server at :%v", httpServer.port)
	http.ListenAndServe(":"+httpServer.port, nil)
}

func (httpServer *HttpServer) LoginAttempt(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	var command struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	if !getBodyAndUnmarshal(w, r, &command) {
		return
	}

	loginSuccess, err := httpServer.storage.CheckLogin(command.Login, command.Password)

	if err != nil {
		log.Printf("error checking DB password:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	type result struct {
		Ok         bool   `json:"ok"`
		StatusText string `json:"statusText"`
	}
	if loginSuccess {
		resultOK := result{
			Ok:         true,
			StatusText: "",
		}
		httpServer.marshalAndSendJson(w, resultOK)
	} else {
		resultNotOk := result{
			Ok:         false,
			StatusText: "failed",
		}
		httpServer.marshalAndSendJson(w, resultNotOk)
	}
}

func (httpServer *HttpServer) Register(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var command struct {
		Login    string `json:"login"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if !getBodyAndUnmarshal(w, r, &command) {
		return
	}
	err := httpServer.storage.Register(command.Login, command.Password)
	if err != nil {
		log.Printf("error checking DB password: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	login := command.Login
	if len(login) > 8 {
		login = login[:8]
	}
	type result struct {
		Login string
	}
	res := result{
		Login: login,
	}
	httpServer.marshalAndSendJson(w, res)
}

func (httpServer *HttpServer) ReturnLogin(w http.ResponseWriter, r *http.Request) {

	login, err := httpServer.storage.FindLogin()
	if err != nil {
		log.Printf("error fetching logged in user DB login: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	type result struct {
		Login string `json:"login"`
		Found bool   `json:"found"`
	}
	if len(login) > 8 {
		login = login[:8]
	}
	found := !(login == "")
	res := result{
		Login: login,
		Found: found,
	}
	httpServer.marshalAndSendJson(w, res)
}

func (HttpServer *HttpServer) Logout(w http.ResponseWriter, r *http.Request) {
	err := HttpServer.storage.Logout()
	if err != nil {
		log.Printf("logout DB error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	w.WriteHeader(http.StatusOK)
}

func (httpServer *HttpServer) marshalAndSendJson(w http.ResponseWriter, data interface{}) {
	jsAnswer, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshaling answer: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	httpServer.sendJson(w, jsAnswer)
}

func (httpServer *HttpServer) sendJson(w http.ResponseWriter, json []byte) {
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
}

func getBodyAndUnmarshal(w http.ResponseWriter, r *http.Request, command interface{}) bool {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Cannot read body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}

	err = json.Unmarshal(body, command)
	if err != nil {
		log.Printf("Unmarshalling body error: %v. Body: %v", err, string(body))
		w.WriteHeader(http.StatusBadRequest)
		return false
	}

	return true
}
