package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// API, API ana sayfasını temsil eder
type API struct {
	Message string `json:"message"`
}

// User, kullanıcı bilgilerini temsil eder
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	apiRoot := "/api"

	// .../api
	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := API{"API Home"}
		output, err := json.Marshal(message)
		checkError(err)
		//setJSONContentType(w)
		fmt.Fprintf(w, string(output))
	})

	// .../api/me
	http.HandleFunc(apiRoot+"/me", func(w http.ResponseWriter, r *http.Request) {
		user := User{ID: 3, FirstName: "Said", LastName: "Darıyemez", Age: 20}
		message := user
		output, err := json.Marshal(message)
		checkError(err)
		//setJSONContentType(w)
		fmt.Fprintf(w, string(output))
	})

	// .../api/users
	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			{ID: 1, FirstName: "Murtaza", LastName: "Kılkuyruk", Age: 77},
			{ID: 2, FirstName: "Hakkı", LastName: "Bulut", Age: 67},
			{ID: 3, FirstName: "Müslüm", LastName: "Gürses", Age: 80},
		}
		message := users
		output, err := json.Marshal(message)
		checkError(err)
		////setJSONContentType(w)
		fmt.Fprintf(w, string(output))
	})

	// 8080 portunda hizmeti başlat
	fmt.Println("Server is listening on :8080...")
	http.ListenAndServe(":8080", nil)
}

// Hata kontrolü yapar ve gerektiğinde programı durdurur
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error:", err.Error())
		os.Exit(1)
	}
}

/*
// İstek yanıtlarının JSON içeriğini ayarlar
func setJSONContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
*/
