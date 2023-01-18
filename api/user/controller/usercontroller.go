package user_controller

import (
	"encoding/json"
	"net/http"

	model "user-api-rest/database-postgresql/model"
	user_repository "user-api-rest/database-postgresql/repository"

	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	router.HandleFunc("/users", getUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", getUserHandler).Methods("GET")
	router.HandleFunc("/users", postUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", deleteUserHandler).Methods("DELETE")
}

func getUsersHandler(responseWritter http.ResponseWriter, request *http.Request) {
	users := user_repository.FindAll()
	responseWritter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWritter).Encode(&users)

}

func getUserHandler(responseWritter http.ResponseWriter, request *http.Request) {
	pathParameters := mux.Vars(request)
	user := user_repository.FindById(pathParameters["id"])
	
	if user.ID == 0 {
		responseWritter.WriteHeader(http.StatusNotFound)
		return
	} 
	
	responseWritter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWritter).Encode(&user)
}

func postUserHandler(responseWritter http.ResponseWriter, request *http.Request) {
	var user model.User
	json.NewDecoder(request.Body).Decode(&user)
	error := user_repository.Save(user)
	if error != nil {
		responseWritter.WriteHeader(http.StatusInternalServerError)
		return 
	}
	responseWritter.WriteHeader(http.StatusCreated)
	json.NewEncoder(responseWritter).Encode(&user)
}

func deleteUserHandler(responseWritter http.ResponseWriter, request *http.Request) {
	pathParameters := mux.Vars(request)
	error := user_repository.Delete(pathParameters["id"])
	if error != nil {
		responseWritter.WriteHeader(http.StatusNotFound)
		return
	}
	responseWritter.WriteHeader(http.StatusOK)
}

