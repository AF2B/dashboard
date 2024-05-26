package routes

import (
	"dashboard/internal/handlers"
	"dashboard/log"
	"fmt"
	"net/http"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type Router struct {
	Routes []*Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) AddRoute(method, path string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	r.Routes = append(r.Routes, &Route{Method: method, Path: path, Handler: http.HandlerFunc(handlerFunc)})
}

func (r *Router) DisplayRoutes() {
	fmt.Println("Rotas atuais:")
	for _, route := range r.Routes {
		fmt.Printf("Método: %s, Caminho: %s\n", route.Method, route.Path)
	}
}

func (r *Router) ConfigureRoutes() {
	log.LogInfo("Configurando rotas para a aplicação")
	r.configureActivityVisitorsRoutes()
	r.configurePostRoutes()
	r.DisplayRoutes()
}

func (r *Router) RegisterRoutes() {
	for _, route := range r.Routes {
		http.HandleFunc(route.Path, route.Handler)
	}
}

func (r *Router) configureActivityVisitorsRoutes() {
	r.AddRoute("GET", "/api/activity-visitors", handlers.ActivityVisitors)
}

func (r *Router) configurePostRoutes() {
	r.AddRoute("GET", "/api/visitor-per-minute", handlers.VisitorPerMinute)
}
