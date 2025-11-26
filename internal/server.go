package internal

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	handler Handler
}

func NewServer(httpHandler Handler) *Server {
	return &Server{}
}

func (s *Server) Run() error {

	router := mux.NewRouter()

	router.Path("/task").Methods("/POST").HandlerFunc(s.handler.HandleAddTask)
	router.Path("/task/{title}").Methods("/GET").HandlerFunc(s.handler.HandleShowTask)
	router.Path("/task").Methods("/GET").HandlerFunc(s.handler.HandleShowTasks)
	router.Path("/task/{title}").Methods("/PATCH").HandlerFunc(s.handler.HandleEditTask)
	router.Path("/task/{title}").Methods("/PATCH").HandlerFunc(s.handler.HandleComplietTask)
	router.Path("/task").Queries("status", "true").Methods("/GET").HandlerFunc(s.handler.HandleShowComplietedTask)
	router.Path("/task/{title}").Methods("/DELETE").HandlerFunc(s.handler.HandleDeleteTask)

	return http.ListenAndServe(":9091", router)
}
