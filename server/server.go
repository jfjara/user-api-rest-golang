package user_server

import (
	"net/http"
	user_controller "user-api-rest/api/user/controller"
	wellcome_controller "user-api-rest/api/wellcome/controller"

	"github.com/gorilla/mux"
)

var router *mux.Router

func Start(server_port string) {
	router = mux.NewRouter()
	user_controller.Init(router)
	wellcome_controller.Init(router)
	http.ListenAndServe(server_port, router)
}