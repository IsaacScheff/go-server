package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
    "os"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", handlePage)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // default port if not specified
    }

	//const addr = ":8080"
	srv := http.Server{
		Handler:      m,
		Addr:         port,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	// this blocks forever, until the server
	// has an unrecoverable error
	fmt.Println("server started on ", addr)
	err := srv.ListenAndServe()
	log.Fatal(err)

    os.Getenv()
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	const page = `<html>
<head></head>
<body>
	<p> Hello from Docker! I'm a Go server. </p>
</body>
</html>
`
	w.Write([]byte(page))
}
