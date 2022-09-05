package main

import (
	"fmt"
	"net/http"

	"translate/controller/api"
)

func main() {

	http.HandleFunc("/getalllanguages", api.GetAllLanguagesFromGoogleTranslate)
	http.HandleFunc("/translate", api.TranslateTheText)

	err := http.ListenAndServe(":8080", nil)
	go func() {
		if err == nil {
			fmt.Println("server started on port 8080", err)
		}
	}()

}
