// handlers/studentHandler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"fealtyx-student-api/fealtyx-student-api/utils"
	"fealtyx-student-api/fealtyx-student-api/models"

	"github.com/gorilla/mux"
)

// CreateStudent - POST /students
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	models.Students[student.ID] = student
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

// GetAllStudents - GET /students
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students := make([]models.Student, 0, len(models.Students))
	for _, student := range models.Students {
		students = append(students, student)
	}
	json.NewEncoder(w).Encode(students)
}

// GetStudentByID - GET /students/{id}
func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	student, exists := models.Students[id]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(student)
}

// UpdateStudentByID - PUT /students/{id}
func UpdateStudentByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var updatedStudent models.Student
	if err := json.NewDecoder(r.Body).Decode(&updatedStudent); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	models.Students[id] = updatedStudent
	json.NewEncoder(w).Encode(updatedStudent)
}

// DeleteStudentByID - DELETE /students/{id}
func DeleteStudentByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if _, exists := models.Students[id]; !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	delete(models.Students, id)
	w.WriteHeader(http.StatusNoContent)
}

// GetStudentSummary - GET /students/{id}/summary
func GetStudentSummary(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	student, exists := models.Students[id]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	summary, err := utils.GetStudentSummary(student)
	if err != nil {
		http.Error(w, "Failed to get summary", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"summary": summary})
}
