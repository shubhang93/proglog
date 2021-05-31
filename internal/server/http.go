package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type httpServer struct {
	Log *Log
}

func NewHTTPServer(addr string) *http.Server {
	r := mux.NewRouter()
	return &http.Server{Addr: addr, Handler: r}
}

func newHTTPServer() *httpServer {
	return &httpServer{
		Log: NewLog(),
	}
}

type ProduceRequest struct {
	Record Record `json:"record"`
}

type ProduceResponse struct {
	Offset uint64 `json:"offset"`
}

type ConsumeRequest struct {
	Offset uint64 `json:"offset"`
}

type ConsumeResponse struct {
	Record Record `json:"record"`
}
