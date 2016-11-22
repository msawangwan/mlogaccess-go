package main

import (
	"fmt"
	"github.com/msawangwan/mgeoloc"
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

	http.HandleFunc("/", indexhandler)
	http.HandleFunc("/favicon.ico", faviconhandler)
	http.HandleFunc("/tyranny", tyrannyhandler)
}

func indexhandler(w http.ResponseWriter, r *http.Request) {
	logaccess(fnewlogentry("index.html", parseip(r)))
	s.pages["index"].ExecuteTemplate(w, "layout", "")
}

func faviconhandler(w http.ResponseWriter, r *http.Request) {
	/* empty handler handles the browser requesting favicon.ico, prevents duplicate log entries */
}

func tyrannyhandler(w http.ResponseWriter, r *http.Request) {
	logaccess(fnewlogentry("tyranny.html", parseip(r)))
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

func fnewlogentry(resource, ip string) string {
	data := mgeoloc.FromAddr(ip)
	return fmt.Sprintf(
		": served %s\n%s",
		resource,
		data.FormatData(),
	)
}

func parseip(r *http.Request) string {
	return r.Header.Get("x-forwarded-for") // change if not using a reverse proxy
}

func main() {
	log.Println("logging access...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}
