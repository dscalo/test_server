package app

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

var dirname, _ = os.Getwd()

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	CORSEnabledFunction(w, r)

	if r.Method != "POST" {
		http.Error(w, r.Method+" not supported", http.StatusBadRequest)
		return
	}

	// set max file size
	err := r.ParseMultipartForm(1024 << 20)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("upload")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	path := path.Join(dirname, "/uploads/", handler.Filename)

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error saving file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer f.Close()

	io.Copy(f, file)

	log.Printf("Successfully ploaded %s \n", handler.Filename)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`success`))

}
