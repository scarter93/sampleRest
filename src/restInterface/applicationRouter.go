/*
* Author: SCARTER
* Date: 04/01/2017
 */
package restInterface

// imports
import (
	"net/http"

	"github.com/gorilla/mux"
)

type ApplicationRoute struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

type ApplicationRoutes []ApplicationRoute

//var router mux.Router
var routes ApplicationRoutes

func CreateApplicationRoute(handlerFunc http.HandlerFunc, method, path string) *ApplicationRoute {
	route := new(ApplicationRoute)
	route.Method = method
	route.Path = path
	route.HandlerFunc = handlerFunc

	return route
}

func BuildApplicationRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Path).Handler(route.HandlerFunc)

	}
	return router
}

func RegisterRoute(route ApplicationRoute) {
	routes = append(routes, route)
}
