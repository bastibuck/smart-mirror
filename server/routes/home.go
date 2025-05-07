package routes

import (
	"net/http"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello World!"))
}
