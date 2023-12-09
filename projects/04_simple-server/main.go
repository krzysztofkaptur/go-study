package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", handlePage)

	const addr = ":8000"

	srv := http.Server{
		Handler: m,
		Addr: addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout: 30 * time.Second,
	}

	fmt.Println("server started on port ", addr)

	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	const page = `
		<html>
			<body>
				<h1>Herro</h1>
			</body>
		</html>
	`

	w.WriteHeader(200)
	w.Write([]byte(page))
}