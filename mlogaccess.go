package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type templates map[string]*template.Template

type site struct {
	pages templates
}

var s *site
var logfile string = "accesslog"

func init() {
	s = &site{
		pages: make(templates),
	}

	s.pages["index"] = template.Must(template.ParseFiles("./template/layout.html", "./template/index.html"))
	s.pages["tyranny"] = template.Must(template.ParseFiles("./template/layout.html", "./template/tyranny.html"))
}

func indexhandler(w http.ResponseWriter, r *http.Request) {
	logaccess(flogentry("index.html", parseip(r)))
	s.pages["index"].ExecuteTemplate(w, "layout", "")
}

func tyrannyhandler(w http.ResponseWriter, r *http.Request) {
	logaccess(flogentry("tyranny.html", parseip(r)))
	s.pages["tyranny"].ExecuteTemplate(w, "layout", "")
}

func logaccess(logentry string) {
	log.Println(logentry)

	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("error opening file: %s", err)
		return
	}

	defer f.Close()

	tstampentry := fmt.Sprintf(
		"%s : %s\n",
		time.Now().Format(time.UnixDate),
		logentry,
	)

	if _, err = f.WriteString(tstampentry); err != nil {
		log.Printf("error writing to file: %s", err)
		return
	}
}

func parseip(r *http.Request) string {
	return r.Header.Get("x-forwarded-for")
}

func flogentry(resource, ip string) string {
	return fmt.Sprintf(
		"[ %s requested %s ]",
		ip,
		resource,
	)
}

func main() {
	log.Println("logging access...")

	http.HandleFunc("/", indexhandler)
	http.HandleFunc("/tyranny", tyrannyhandler)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
