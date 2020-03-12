package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"regexp"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func getPostgreSQLVersion() (string, error) {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return "", err
	}

	rows := []string{}
	err = db.Select(&rows, "SELECT version()")
	if err != nil {
		return "", err
	}

	return rows[0], nil
}

func helloHtml(path string) string {
	no_foo := convertFoo(path)
	return fmt.Sprintf("<h1>Welcome to revAMPD</h1>\n<p>You've requested: %s</p>\n", html.EscapeString(no_foo))
}

func convertFoo(path string) string {
	re := regexp.MustCompile(`(?i)foo`)
	path = re.ReplaceAllString(path, "bar")
	return path
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, helloHtml(r.URL.Path))
	})

	version, err := getPostgreSQLVersion()
	if err != nil {
		log.Error(err)
	} else {
		log.Debug(version)
	}

	log.Debug("Starting revAMPD API backend, listening on port " + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
