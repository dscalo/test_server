package app

import (
	"encoding/json"
	"net/http"
)

type TestResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}
type TestData struct {
	Name string `json:"name"`
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, r.Method+" not supported", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`PONG`))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "NOT FOUND", http.StatusNotFound)
	return
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	CORSEnabledFunction(w, r)
	var data TestData
	switch r.Method {
	case "GET":
		query := r.URL.Query()
		name := query.Get("name")
		if len(name) == 0 {
			http.Error(w, "No name given", http.StatusBadRequest)
			return
		}
		data = TestData{Name: name}
	case "POST":
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, r.Method+" not supported", http.StatusBadRequest)
		return
	}

	testRes := TestResponse{ID: 4, Message: "Hey There! " + data.Name}
	js, _ := json.Marshal(testRes)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
