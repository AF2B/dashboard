package main

import (
	router "dashboard/internal/routes"
	"net/http"
)

func main() {
    r := router.NewRouter()
    r.ConfigureRoutes()
    r.RegisterRoutes()

    http.ListenAndServe(":8080", nil)
}
