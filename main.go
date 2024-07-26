package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerMainPage)
	http.HandleFunc("/get_request", handlerGetRequest)
	fmt.Println("Listening on port 8080: http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func handlerMainPage(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello there"))

	if err != nil {
		log.Fatal(err)
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
	// fmt.Println(r.Body)
	// fmt.Println(r.Form)
	// fmt.Println(r.PostForm)
	_, err := w.Write([]byte("Request Handler"))

	if err != nil {
		log.Fatal(err)
	}

	// ret = strings.Replace(ret, "\n", "<br>", -1)

	_, err = w.Write([]byte(ret))

	if err != nil {
		log.Fatal(err)
	}
}
