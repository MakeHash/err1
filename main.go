package inv

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Storage struct {
	F string `json:"first"`
	S int    `json:"second"`
	// Add more fields as per your JSON data
}

func myServ(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var data Storage

		// Parse JSON from the request body to the struct
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func main() {
	http.HandleFunc("https://err1.vercel.app/api", myServ)
	err := http.ListenAndServe("https://err1.vercel.app:8443", nil)
	if err != nil {
		fmt.Printf("error: %s", err)
	}

}
