package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"simple_server/helper"
)

func main() {
	http.HandleFunc("/", handlerMainPage)
	http.HandleFunc("/get_request", handlerGetRequest)
	http.HandleFunc("/query_with_params", handleQueryWithParams)
	http.HandleFunc("/post_form", handleFormRequest)
	http.HandleFunc("/send_form", handlerPostRequest)
	fmt.Println("Listening on port 8080: http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func handlerMainPage(w http.ResponseWriter, r *http.Request) {
	err := helper.OpenTemplate(w, "index")

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

	err := helper.OpenTemplate(w, "get_request")

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
	err := helper.OpenTemplate(w, "params")

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
	err := helper.OpenTemplate(w, "form")

	if err != nil {
		log.Println(err)
	}
}

func handlerPostRequest(w http.ResponseWriter, r *http.Request) {
	err := helper.OpenTemplate(w, "form_res")

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
