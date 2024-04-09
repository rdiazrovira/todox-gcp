package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"todox/internal"

	"github.com/leapkit/core/envor"
	"github.com/leapkit/core/server"
)

func main() {
	host := envor.Get("HOST", "0.0.0.0")
	port := envor.Get("PORT", "3000")

	slog.Info("Checking 2...")
	slog.Info("> Starting server...")
	server := server.New(
		server.WithHost(host),
		server.WithPort(port),
	)

	// Application services
	slog.Info("> Adding Services...")
	if err := internal.AddServices(server); err != nil {
		fmt.Println("failed while adding services: ", err.Error())
		slog.Error(err.Error())
		os.Exit(1)
	}

	// Application routes
	slog.Info("> Setting up routes...")
	if err := internal.AddRoutes(server); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	slog.Info(fmt.Sprintf("> Starting todox at %s", server.Addr()))
	if err := http.ListenAndServe(server.Addr(), server.Handler()); err != nil {
		slog.Error(fmt.Sprintf("Server terminated: %v", err.Error()))
	}
}
