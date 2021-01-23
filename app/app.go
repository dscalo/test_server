package app

import (
	"log"
	"net/http"
)

func Run() {
	endPoints := map[string]http.HandlerFunc{
		"/":       NotFoundHandler,
		"/ping":   PingHandler,
		"/test":   TestHandler,
		"/upload": UploadHandler,
	}

	for endPoint, fn := range endPoints {
		http.HandleFunc(endPoint, fn)
	}

	log.Fatal(http.ListenAndServe(":3333", nil))

}
