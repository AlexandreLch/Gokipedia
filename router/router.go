package router

import (
	"github.com/gorilla/mux"
	"gokipedia/controllers"
	"net/http"
)

// Route struct defining all of this project routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes slice of Route
type Routes []Route

// newRouter registers public routes
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		Name:        "Say Hello",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: controllers.RenderHome,
	},
	Route{
		Name:        "Get all articles",
		Method:      "GET",
		Pattern:     "/articles",
		HandlerFunc: controllers.RenderArticles,
	},
	Route{
		Name:        "Add an article",
		Method:      "GET",
		Pattern:     "/articles/new",
		HandlerFunc: controllers.RenderArticleForm,
	},
	Route{
		Name:        "Add an article",
		Method:      "POST",
		Pattern:     "/articles/new",
		HandlerFunc: controllers.SaveArticle,
	},
	Route{
		Name:        "Get an article",
		Method:      "GET",
		Pattern:     "/articles/{id}",
		HandlerFunc: controllers.RenderArticle,
	},
	Route{
		Name:        "Export articles",
		Method:      "POST",
		Pattern:     "/articles/export",
		HandlerFunc: controllers.ExportArticles,
	},
}
