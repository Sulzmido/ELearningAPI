package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"./models"

	"github.com/julienschmidt/httprouter"
)

func main() {

	r := httprouter.New()

	r.GET("/", index)

	r.GET("/student/:id", getStudent)
	r.POST("/student", createStudent)
	r.DELETE("/student/:id", deleteStudent)

	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	s := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>E-LEARNING API</title>
</head>
<body>
	WELCOME TO THE E-LEARNING API
</body>
</html>
	`

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getStudent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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

func createStudent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// validate what is in r.Body.
	u := models.Student{}
	json.NewDecoder(r.Body).Decode(&u)
	u.ID = "1"
	// Save u.
	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func deleteStudent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Write code to delete user\n")
}
