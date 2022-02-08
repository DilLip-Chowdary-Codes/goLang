package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Student struct {
	Name  string
	Class string
	About string
}

var Students map[string]Student

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	//Routes
	router.HandleFunc("/", homePage)
	router.HandleFunc("/students", returnAllStudents)
	router.HandleFunc("/student/", createStudent).Methods("POST")
	router.HandleFunc("/student/{id}", returnSingleStudent).Methods("GET")
	router.HandleFunc("/student/{id}", deleteStudent).Methods("DELETE")
	router.HandleFunc("/student/{id}", updateStudent).Methods("PUT")

	fmt.Println("Listening to requests on 8111")
	log.Fatal(http.ListenAndServe(":8111", router))
}

func updateStudent(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	reqBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		fmt.Println("Error: ", err)

	}

	type reqBodyObj struct {
		Student Student
		Details string
	}

	var requestObj reqBodyObj

	err = json.Unmarshal(reqBody, &requestObj)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	studentId := vars["id"]

	if _, exists := Students[studentId]; exists {
		Students[studentId] = requestObj.Student

		writer.WriteHeader(http.StatusCreated)
		_, err := writer.Write([]byte("Updated Student Successfully"))
		if err != nil {
			http.Error(writer, "Some thing Went Wrong", 400)

			return
		}

	} else {
		http.Error(writer, "No Student Found with given ID", 404)
	}
}

func deleteStudent(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: deleteStudent")

	reqVars := mux.Vars(request)

	studentId := reqVars["id"]

	if _, exists := Students[studentId]; exists {
		delete(Students, studentId)
		writer.WriteHeader(http.StatusAccepted)
		_, err := writer.Write([]byte("Deleted Student Successfully"))
		if err != nil {
			http.Error(writer, "Some thing Went Wrong", 400)

			return
		}

	} else {
		http.Error(writer, "No Student Found with given ID", 404)

	}

}

func createStudent(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: createStudent")

	reqBody, _ := ioutil.ReadAll(request.Body)

	type reqBodyObj struct {
		Student Student
		Details string
	}

	var requestObj reqBodyObj

	err := json.Unmarshal(reqBody, &requestObj)

	if err != nil {
		fmt.Println("Error: ", err)
		http.Error(writer, "Wrong Input Given", 400)
		return
	}

	noOfStudents := len(Students)

	Students[uuid.New().String()] = requestObj.Student

	if len(Students) == noOfStudents+1 {
		writer.WriteHeader(http.StatusCreated)
		_, err := writer.Write([]byte("Created Student Successfully"))
		if err != nil {
			fmt.Println("Error: ", err)
			http.Error(writer, "Some thing Went Wrong", 400)
			return
		}
	} else {
		http.Error(writer, "Some thing Went Wrong, Contact Admin", 400)

	}

	if err != nil {
		return
	}
}

func returnSingleStudent(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	err := json.NewEncoder(writer).Encode(Students[id])
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Endpoint Hit: returnSingleStudent")

}

func returnAllStudents(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: returnAllStudents")

	//json.NewEncoder(writer).Encode(Students)
	err := json.NewEncoder(writer).Encode(Students)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write([]byte("Welcome to Homepage 😍"))
	if err != nil {
		http.Error(writer, "Some thing Went Wrong", 400)

		return
	}

	fmt.Println("Endpoint Hit: homePage")
}

func main() {
	Students = map[string]Student{
		uuid.New().String(): {Name: "DilLip", Class: "A", About: "He is Offline Introvert you know 😁"},
		uuid.New().String(): {"Mouni", "A", "She is not DilLip's GF 🙂"},
		uuid.New().String(): {Name: "Raj", Class: "B", About: "Raj is Raj, so Raj is Raj 😉"},
		uuid.New().String(): {Name: "Bob", Class: "C ", About: "You know I Don't know Who is Bob 😅"},
	}

	handleRequests()

}
