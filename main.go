package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*

Simple mockup of a REST API login using the hard coded username and password of
"admin" and "password".
HTML and javascript is embedded in the built binary

*/

//go:embed index.html
var content embed.FS

// Define the Credentials struct here
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/api/login", handleLogin)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	data, err := content.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Could not read index.html", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hardcoded username and password for demonstration
	const validUsername = "admin"
	const validPassword = "password"

	// Check if the provided credentials match the hardcoded ones
	if creds.Username == validUsername && creds.Password == validPassword {
		// Successful login
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Logged in successfully as %s", creds.Username)
	} else {
		// Unsuccessful login
		http.Error(w, "Unauthorized: Invalid credentials", http.StatusUnauthorized)
	}
}
