package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type email struct {
	To      string
	From    string
	Subject string
	Date    string
}

var emails []email

func GetEmailsData(w http.ResponseWriter, req *http.Request) {
	client := http.Client{}
	user := "ogsystem@gmail.com"
	pass := "Complexpass#123"
	url := "http://localhost:5080/api/default/_search"

	query := `{"query": {"from": 0, "size": 1000, "sql": "SELECT * FROM \"email\" WHERE id is not null"}}`

	r, err := http.NewRequest("POST", url, strings.NewReader(query))
	if err != nil {
		log.Println("No fue posible conectarse ", err)
	} else {
		r.SetBasicAuth(user, pass)
		response, err := client.Do(r)
		if err != nil {
			log.Println("Sin respuesta ", err)
		} else {
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			} else {
				mapResponse := make(map[string]interface{})
				json.Unmarshal(body, &mapResponse)
				json.NewEncoder(w).Encode(mapResponse)
			}
			defer response.Body.Close()
		}
	}
}

func main() {
	router := chi.NewRouter()

	router.Get("/emails", GetEmailsData)

	log.Println(http.ListenAndServe(":3000", router))
}
