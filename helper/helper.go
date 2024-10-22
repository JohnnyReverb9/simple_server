package helper

import (
	"embed"
	"log"
	"net/http"
	"strings"
	"time"
)

type TimeHandler struct {
	Format string
}

func (th *TimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.Format)
	_, err := w.Write([]byte("The time is: " + tm))

	if err != nil {
		log.Println(err)
		return
	}
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		_, err := w.Write([]byte("GET"))

		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		_, err := w.Write([]byte("POST"))

		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPut:
		_, err := w.Write([]byte("PUT"))

		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPatch:
		_, err := w.Write([]byte("PATCH"))

		if err != nil {
			log.Fatal(err)
		}
	case http.MethodDelete:
		_, err := w.Write([]byte("DELETE"))

		if err != nil {
			log.Fatal(err)
		}
	}
}

func OpenTemplate(w http.ResponseWriter, title string, static embed.FS) error {
	html, err := static.ReadFile("html/" + title + ".html")

	if err != nil {
		http.Error(w, "Could not read HTML file", http.StatusInternalServerError)
		panic(err)
	}

	_, err = w.Write(html)

	if err != nil {
		panic(err)
	}

	return nil
}

func NToBrReplacer(str string) string {
	return strings.Replace(str, "\n", "<br>", -1)
}
