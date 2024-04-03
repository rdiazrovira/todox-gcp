package internal

import (
	"fmt"
	"todox/internal/todos"

	"github.com/leapkit/core/server"

	lserver "github.com/leapkit/core/server"
)

// AddServices is a function that will be called by the server
// to inject services in the context.
func AddServices(r server.Router) error {
	conn, err := DB()
	if err != nil {
		fmt.Println("Failed to start database connection: ", err.Error())
		return err
	}

	// Services that will be injected in the context
	r.Use(lserver.InCtxMiddleware("todoService", todos.NewService(conn)))

	return nil
}
