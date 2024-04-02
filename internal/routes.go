package internal

import (
	"embed"
	"todox/internal/todos"
	"todox/public"

	"github.com/leapkit/core/envor"
	"github.com/leapkit/core/render"
	"github.com/leapkit/core/server"
	"github.com/leapkit/core/session"
)

//go:embed **/*.html *.html
var tmpls embed.FS

// AddRoutes mounts the routes for the application,
// it assumes that the base services have been injected
// in the creation of the server instance.
func AddRoutes(r server.Router) error {
	// Session middleware to be used by the application
	// to store session data.
	r.Use(session.Middleware(
		envor.Get("SESSION_SECRET", "secret_key"),
		envor.Get("SESSION_NAME", "todox_session"),
	))

	// Render middleware to build the html templates
	// and serve them to the client.
	r.Use(render.Middleware(
		tmpls,
		render.WithDefaultLayout("layout.html")),
	)

	r.HandleFunc("GET /{$}", todos.Index)
	r.HandleFunc("GET /search", todos.Search)
	r.HandleFunc("POST /{$}", todos.Create)

	r.Group("/{id}/", func(wid server.Router) {
		wid.HandleFunc("GET /edit", todos.Edit)
		wid.HandleFunc("GET /show", todos.Show)
		wid.HandleFunc("DELETE /{$}", todos.Delete)
		wid.HandleFunc("PUT /{$}", todos.Update)
		wid.HandleFunc("PUT /complete", todos.Complete)
	})

	// Mount the public folder to be served openly
	r.Folder("/public", public.Files)

	return nil
}
