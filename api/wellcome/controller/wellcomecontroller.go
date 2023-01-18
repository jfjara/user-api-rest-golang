package wellcome_controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func wellcomePageHandler(responseWritter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWritter, "Wellcome page\n")
	responseWritter.Write([]byte("This is a wellcome test page"))
}

func Init(router *mux.Router) {
	router.HandleFunc("/", wellcomePageHandler)
}