package helper

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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

func OpenTemplate(w http.ResponseWriter, title string) error {
	html, err := ioutil.ReadFile("./html/" + title + ".html")

	if err != nil {
		http.Error(w, "Could not read HTML file", http.StatusInternalServerError)
		return err
	}

	_, err = w.Write(html)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func NToBrReplacer(str string) string {
	return strings.Replace(str, "\n", "<br>", -1)
}
