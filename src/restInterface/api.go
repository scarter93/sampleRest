package restInterface

// imports
import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"sampleRest/src/data"
)

var related []int
var sampleTask = data.CreateNewTask(9999999, time.Date(9999, time.December, 31, 12, 0, 0, 0, time.Local), &related, "Sample Task", "This is a sample task for testing purposes")
var tasksMap map[int]data.Task = make(map[int]data.Task)

// Base function
func Base(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, you have reached the base page: %q", html.EscapeString(request.URL.Path))
}

func GetTaskById(writer http.ResponseWriter, request *http.Request) {
	var taskId int
	var err error
	var task data.Task

	variables := mux.Vars(request)
	if taskId, err = strconv.Atoi(variables["taskId"]); err != nil {
		panic(err)
	}
	//get task
	task = tasksMap[taskId]

	//if we did not find the task
	if task.Id <= 0 {
		setContentType(writer, "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	//if we found the data display
	setContentType(writer, "application/json; charset=UTF-8")
	json.NewEncoder(writer).Encode(task)
}

func GetSampleTask(writer http.ResponseWriter, request *http.Request) {
	setContentType(writer, "application/json; charset=UTF-8")
	json.NewEncoder(writer).Encode(sampleTask)
}

func SaveSampleTask(writer http.ResponseWriter, request *http.Request) {
	var task data.Task
	jsonBytes, err := ioutil.ReadAll(io.LimitReader(request.Body, 10485760))
	processErrors(writer, request, err)
	if err := json.Unmarshal(jsonBytes, &task); err != nil {
		writer.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(writer).Encode(err); err != nil {
			panic(err)
		}
	}
	setContentType(writer, "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusCreated)
	id := task.Id
	tasksMap[id] = task
	json.NewEncoder(writer).Encode(task)
}

//Helpers

func setContentType(writer http.ResponseWriter, contentType string) {
	writer.Header().Set("Content-Type", contentType)
}

func processErrors(writer http.ResponseWriter, request *http.Request, err error) {
	if err != nil {
		log.Panic(err)
		panic(err)
	}
	if err := request.Body.Close(); err != nil {
		log.Panic(err)
		panic(err)

	}
}
