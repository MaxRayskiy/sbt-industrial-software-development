package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type Issue struct {
	Title       string
	Description string
	Email       string
	IsPrivate   bool
}

func main() {
	// some time to
	time.Sleep(10)
	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	database, err := Initialize(dbUser, dbPassword, dbName)

	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer database.Conn.Close()

	http.HandleFunc("/add", database.Add)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (db *Database) Add(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("issue.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := Issue{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Email:       r.FormValue("email"),
		IsPrivate:   true,
	}

	query := `INSERT INTO issues (title, description, email, is_private) VALUES ($1, $2, $3, $4)`
	err := db.Conn.QueryRow(query, details.Title, details.Description, details.Email, details.IsPrivate)
	if err != nil {
		log.Println(err.Scan())
	}

	tmpl.Execute(w, struct{ Success bool }{true})
}
