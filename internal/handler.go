package internal

import "net/http"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// /task
// /POST
// 201 400 500
func (h *Handler) HandleAddTask(w http.ResponseWriter, r *http.Request) {

}

// /task/{title}
// /GET
// 200 400 500
func (h *Handler) HandleShowTask(w http.ResponseWriter, r *http.Request) {

}

// /task
// /GET
// 200 400 500
func (h *Handler) HandleShowTasks(w http.ResponseWriter, r *http.Request) {

}

// /task/{title}
// /PATCH
// 200 400 500
func (h *Handler) HandleEditTask(w http.ResponseWriter, r *http.Request) {

}

// task/{title}
// /PATCH
// 200 400 500
func (h *Handler) HandleComplietTask(w http.ResponseWriter, r *http.Request) {

}

// /task/?status=true
// /PATCH
//200 400 500
func (h *Handler) HandleShowComplietedTask(w http.ResponseWriter, r *http.Request) {

}

// task/{title}
// /DELETE
// 200 400 500
func (h *Handler) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {

}
