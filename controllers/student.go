package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"github.com/julienschmidt/httprouter"
)

// StudentController struct
type StudentController struct {
}

// NewStudentController constructor
func NewStudentController() *StudentController {
	return &StudentController{}
}

// GetStudent GET
func (c StudentController) GetStudent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	student := models.Student{
		ID:     p.ByName("id"),
		Name:   "Sulaiman Ahmed",
		Gender: "Male ",
		Age:    21,
	}

	uj, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

// CreateStudent POST
func (c StudentController) CreateStudent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// validate what is in r.Body.
	student := models.Student{}
	json.NewDecoder(r.Body).Decode(&student)
	student.ID = "1"
	// Save u.
	uj, _ := json.Marshal(student)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

// DeleteStudent DELETE
func (c StudentController) DeleteStudent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Write code to delete user\n")
}
