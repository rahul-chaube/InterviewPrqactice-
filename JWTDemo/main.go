package main

import (
	"JWTDemo/auth"
	"JWTDemo/utility"
	"fmt"
	"net/http"
)

func main() {

	http.Handle("/create", auth.Logging(http.HandlerFunc(createToken)))
	http.Handle("/get", auth.Logging(auth.AuthMiddleware(http.HandlerFunc(AddUser))))

	http.ListenAndServe(":8088", nil)
}

func createToken(res http.ResponseWriter, req *http.Request) {

	query := req.URL.Query()
	username := query.Get("username")
	if username == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("username Not found "))
		return
	}

	tokken, err := utility.CreateToken(username)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Error occured  " + err.Error()))
		return
	}

	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(tokken))

}

func AddUser(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Method Not allowed "+req.Method, http.StatusMethodNotAllowed)
		return
	}
	username := req.Context().Value("username")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(fmt.Sprintf("Hello %s ", username)))

}
