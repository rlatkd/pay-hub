package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	"auth/internal/handlers"
	"auth/internal/repository"
	"auth/internal/utils"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	dsn := "rlatkd:0000@tcp(127.0.0.1:3300)/AUTH?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("DB 연결 실패: %v\n", err)
	}
	fmt.Println("MySQL Connected (localhost:3300)")

	hash, _ := bcrypt.GenerateFromPassword([]byte("0000"), bcrypt.DefaultCost)

	repository.CreateTestUser(db, string(hash))

	fmt.Println("Test User ensured: admin / 0000")

	h := &handlers.AuthHandler{DB: db}

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.LoginPage(w, r)
		} else if r.Method == http.MethodPost {
			h.LoginProcess(w, r)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusFound)
	})

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Auth Server started at http://localhost:8080")

	go utils.Opener("http://localhost:8080")

	log.Fatal(http.Serve(listener, nil))
}
