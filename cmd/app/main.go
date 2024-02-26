package main

import (
	"fmt"
	"log/slog"
	"os"
	"todox/internal"

	"github.com/leapkit/core/server"
)

func main() {
	server := server.New(
		server.WithHost(internal.Host),
		server.WithPort(internal.Port),
	)

	// Application services
	if err := internal.AddServices(server); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// Application routes
	if err := internal.AddRoutes(server); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := server.Start(); err != nil {
		slog.Error(fmt.Sprintf("Server terminated: %v", err.Error()))
	}
}
