package main

import (
	"fmt"
	"net/http"

	"github.com/MattBabbage/goRestAPI/internal/routes"
)

func Hello() string {
	return "Hello, world"
}

func main() {
	router := routes.APIRouter()
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server Opened on port %s", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
