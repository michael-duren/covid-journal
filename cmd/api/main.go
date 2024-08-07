package main

import (
	"covid-journal/internal/auth"
	"covid-journal/internal/server"
	"fmt"
)

func main() {
	auth.NewAuth()
	server := server.NewServer()

	fmt.Println("Server listening on port:", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
