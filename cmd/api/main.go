package main

import (
	"covid-journal/internal/auth"
	"covid-journal/internal/server"
	"fmt"
	"os"
)

func main() {
	auth.NewAuth()
	server := server.NewServer()
	env := os.Getenv("APP_ENV")

	if env == "local" {
		fmt.Printf("Server running at http://localhost%s\n", server.Addr)
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
