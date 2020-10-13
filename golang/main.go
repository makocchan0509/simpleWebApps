package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	startWebServer("8080")
}

func startWebServer(port string) error {

	log.Println("Running web server ...")
	http.HandleFunc("/", apiMakerHandler(simpleHandler))
	return http.ListenAndServe(":"+port, nil)
}

func apiMakerHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

type APIResponse struct {
	Result     string `json:"result"`
	ErrMessage string `json:"errMessage"`
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {

	var response APIResponse

	response = APIResponse{
		Result:     "Success",
		ErrMessage: "Hello Golang World!",
	}
	output, err := json.Marshal(response)
	if err != nil {
		log.Println("failed json marchal", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
