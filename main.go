// sampleRest project main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"sampleRest/src/restInterface"
)

func main() {
	fmt.Println("Starting sample REST service")

	route := *restInterface.CreateApplicationRoute(restInterface.Base, "GET", "/")
	restInterface.RegisterRoute(route)
	route = *restInterface.CreateApplicationRoute(restInterface.GetTaskById, "GET", "/task/{taskId}")
	restInterface.RegisterRoute(route)
	route = *restInterface.CreateApplicationRoute(restInterface.GetSampleTask, "GET", "/sampleTask")
	restInterface.RegisterRoute(route)
	route = *restInterface.CreateApplicationRoute(restInterface.SaveSampleTask, "POST", "/task")
	restInterface.RegisterRoute(route)

	//Build
	router := restInterface.BuildApplicationRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
