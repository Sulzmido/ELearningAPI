package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	r := httprouter.New()

	r.GET("/", index)

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
