// main.go
package main

import (
	"fealtyx-student-api/fealtyx-student-api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/students", handlers.CreateStudent).Methods("POST")
	router.HandleFunc("/students", handlers.GetAllStudents).Methods("GET")
	router.HandleFunc("/students/{id}", handlers.GetStudentByID).Methods("GET")
	router.HandleFunc("/students/{id}", handlers.UpdateStudentByID).Methods("PUT")
	router.HandleFunc("/students/{id}", handlers.DeleteStudentByID).Methods("DELETE")
	router.HandleFunc("/students/{id}/summary", handlers.GetStudentSummary).Methods("GET")

	http.ListenAndServe(":8080", router)
}
