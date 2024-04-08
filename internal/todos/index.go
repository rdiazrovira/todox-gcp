package todos

import (
	"log/slog"
	"net/http"

	"github.com/leapkit/core/envor"
	"github.com/leapkit/core/render"
)

func Index(w http.ResponseWriter, r *http.Request) {
	host := envor.Get("HOST", "0.0.0.0")
	port := envor.Get("HOST_PORT", "3000")

	slog.Info("> Loading the todos... ", host, port)

	todos := r.Context().Value("todoService").(*service)

	list, err := todos.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rw := render.FromCtx(r.Context())
	rw.Set("list", list)

	err = rw.Render("todos/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
