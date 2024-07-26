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
	err := OpenTemplate(w, "index")

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
	// fmt.Println(r.Body)
	// fmt.Println(r.Form)
	// fmt.Println(r.PostForm)
	// _, err := w.Write([]byte("Request Handler"))

	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := OpenTemplate(w, "get_request")

	if err != nil {
		log.Println(err)
		return
	}

	ret = NToBrReplacer(ret)

	_, err = w.Write([]byte(ret))

	if err != nil {
		log.Fatal(err)
	}
}
