package httpserver

import (
	"log"
	"net/http"
	"os"
	//"github.com/gorilla/mux"
)

const defPort = "81"
const apiPrefix = "/api"

type HttpServer struct {
	port        string
	siteHandler http.Handler
	storage     StorageItf
}

func Init(storage StorageItf) *HttpServer {
	httpServer := &HttpServer{
		storage: storage,
		port:    defPort,
	}

	if os.Getenv("PORT") != "" {
		httpServer.port = os.Getenv("PORT")
	}

	//router := mux.NewRouter()
	//unsecureApi := router.PathPrefix(apiPrefix).Subrouter()
	//unsecureApi.HandleFunc("/send", httpServer.SendSheet).Methods("GET")
	return httpServer
}

func InitAndStart(storage StorageItf) *HttpServer {
	httpServer := Init(storage)
	httpServer.Start()
	return httpServer
}

func (httpServer *HttpServer) Start() {
	log.Printf("Start HTTP Server at :%v", httpServer.port)
	log.Fatalf("HTTP Server ListenAndServe fatal error: %v", http.ListenAndServe(":"+httpServer.port, httpServer.siteHandler))
}
