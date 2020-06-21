package main

import (
	"net/http"

	"./controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {

	homeController := controllers.NewHomeController()
	studentController := controllers.NewStudentController()

	r := httprouter.New()

	r.GET("/", homeController.Index)

	r.GET("/student", studentController.GetStudents)
	r.GET("/student/:id", studentController.GetStudent)
	r.POST("/student", studentController.CreateStudent)
	r.DELETE("/student/:id", studentController.DeleteStudent)

	http.ListenAndServe(":8080", r)
}
