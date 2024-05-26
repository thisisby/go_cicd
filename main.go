package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Tom = Person{
	Id:   1,
	Name: "Tom",
	Age:  20,
}

func TomHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		j, _ := json.Marshal(Tom)
		w.Write(j)
	case http.MethodPost:
		d := json.NewDecoder(r.Body)
		p := &Person{}
		err := d.Decode(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		Tom = *p
		w.Write([]byte("OK"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
func connectDB() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	return sql.Open("postgres", connectionString)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	log.Printf("Connected to the database")
	log.Printf("Postgres version: %v", db.Ping())
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS people (id SERIAL, name VARCHAR, age INTEGER)"); err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	server := http.NewServeMux()
	server.HandleFunc("/", TomHandler)
	server.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			rows, err := db.Query("SELECT name, age FROM people")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			people := []Person{}
			for rows.Next() {
				p := Person{}
				if err := rows.Scan(&p.Name, &p.Age); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				people = append(people, p)
			}

			j, _ := json.MarshalIndent(people, "", "  ")
			w.Write(j)
		case http.MethodPost:
			d := json.NewDecoder(r.Body)
			p := &Person{}
			err := d.Decode(p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if _, err := db.Exec("INSERT INTO people (name, age) VALUES ($1, $2)", p.Name, p.Age); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte("OK"))
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", server)
}
