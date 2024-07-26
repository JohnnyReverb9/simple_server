package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"simple_server/helper"
	"time"
)

//go:embed html/*
var static embed.FS

func main() {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/", handlerMainPage)
	serverMux.HandleFunc("/get_request", handlerGetRequest)
	serverMux.HandleFunc("/query_with_params", handleQueryWithParams)
	serverMux.HandleFunc("/post_form", handleFormRequest)
	serverMux.HandleFunc("/send_form", handlerPostRequest)

	th1123 := &helper.TimeHandler{Format: time.RFC1123}
	serverMux.Handle("/time/1", th1123)

	th3339 := &helper.TimeHandler{Format: time.RFC3339}
	serverMux.Handle("/time/2", th3339)

	fmt.Println("Begin listening on port 8080: http://localhost:8080")

	err := http.ListenAndServe(":8080", serverMux)

	if err != nil {
		log.Fatal(err)
	}
}

func handlerMainPage(w http.ResponseWriter, r *http.Request) {
	err := helper.OpenTemplate(w, "index", static)

	if err != nil {
		log.Println(err)
	}
}

func handlerGetRequest(w http.ResponseWriter, r *http.Request) {
	ret := ""
	fmt.Println(r.Method) // Тип метода
	ret += r.Method + "\n\n"
	fmt.Println(r.URL) // запрашиваемый URL
	ret += r.URL.Path + "\n\n"
	fmt.Println(r.Proto) // версия протокола
	ret += r.Proto + "\n\n"
	fmt.Println(r.Host)
	ret += r.Host + "\n\n"
	fmt.Println(r.UserAgent())
	ret += r.UserAgent() + "\n\n"
	fmt.Println(r.Header)
	ret += fmt.Sprintf("%v", r.Header)

	err := helper.OpenTemplate(w, "get_request", static)

	if err != nil {
		log.Println(err)
		return
	}

	ret = helper.NToBrReplacer(ret)

	_, err = w.Write([]byte(ret))

	if err != nil {
		log.Fatal(err)
	}
}

func handleQueryWithParams(w http.ResponseWriter, r *http.Request) {
	err := helper.OpenTemplate(w, "params", static)

	if err != nil {
		log.Println(err)
	}

	ret := ""

	ret += "URL: " + r.URL.String() + "\n\n"

	if r.URL.Query().Has("action") {
		ret += "action: " + r.URL.Query().Get("action") + "\n\n"
	}

	if r.URL.Query().Has("lang") {
		switch r.URL.Query().Get("lang") {
		case "en":
			ret += "lang: " + r.URL.Query().Get("lang") + "\n\n"
		case "ru":
			ret += "язык: " + r.URL.Query().Get("lang") + "\n\n"
		}
	}

	ret = helper.NToBrReplacer(ret)

	_, err = w.Write([]byte(ret))

	if err != nil {
		log.Fatal(err)
	}
}

func handleFormRequest(w http.ResponseWriter, r *http.Request) {
	err := helper.OpenTemplate(w, "form", static)

	if err != nil {
		log.Println(err)
	}
}

func handlerPostRequest(w http.ResponseWriter, r *http.Request) {
	err := helper.OpenTemplate(w, "form_res", static)

	if err != nil {
		log.Println(err)
	}

	if r.Method == http.MethodPost {
		bytesBody, err := io.ReadAll(r.Body)

		if err != nil {
			log.Println(err)

			_, err = w.Write([]byte("400 | Bad request"))

			if err != nil {
				log.Println(err)
			}

			return
		}

		_, err = w.Write([]byte(helper.NToBrReplacer(string(bytesBody) + "\n\nOK!")))

		if err != nil {
			log.Println(err)
		}
		return
	}

	_, err = w.Write([]byte("POST method available only"))

	if err != nil {
		log.Println(err)
	}
}
