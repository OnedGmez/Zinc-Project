package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

var (
	client = http.Client{}
	query  string
	emails []email
)

const (
	user = "ogsystem@gmail.com"
	pass = "Complexpass#123"
)

/*
PostEmailsData función de endpoint utilizada para enviar los datos a almacenar en openobserve

@Versión: 1.0
*/
func PostEmailsData(w http.ResponseWriter, req *http.Request) {
	url := "http://localhost:5080/api/default/email/_json"
	content, _ := io.ReadAll(req.Body)
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(content))
	if err != nil {
		log.Println("No fue posible conectarse ", err)
	} else {
		r.SetBasicAuth(user, pass)
		response, err := client.Do(r)
		if err != nil {
			fmt.Println("Sin respuesta ", err)
		} else {
			defer response.Body.Close()
		}
	}
}

/*
GetEmailsData función de endpoint utilizada para extraer los datos que necesitamos desde Openobserve
*/
func GetEmailsData(w http.ResponseWriter, req *http.Request) {
	url := "http://localhost:5080/api/default/_search"
	criterion := req.URL.Query()["criterion"]
	value := req.URL.Query()["value"]
	if criterion[0] == "to" || criterion[0] == "subject" {
		query = `{"query": {"from": 0, "size": 80000, "sql": "SELECT * FROM \"email\" WHERE` + ` \"` + criterion[0] + `\" LIKE ` + `'%` + strings.TrimSpace(value[0]) + `%' ORDER BY id ASC"}}`
	} else {
		query = `{"query": {"from": 0, "size": 80000, "sql": "SELECT * FROM \"email\" WHERE` + ` \"from\" LIKE ` + `'` + strings.TrimSpace(value[0]) + `%' ORDER BY id ASC"}}`
	}
	r, err := http.NewRequest("POST", url, strings.NewReader(query))
	if err != nil {
		log.Println("No fue posible conectarse ", err)
	} else {
		r.SetBasicAuth(user, pass)
		w.Header().Add("Access-Control-Allow-Origin", "*")
		response, err := client.Do(r)
		if err != nil {
			log.Println("Sin respuesta ", err)
		} else {
			body, err := io.ReadAll(response.Body)
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

/*
DelEmailsData función de endpoint utilizada para eliminar el stream de datos
*/
func DelEmailsData(w http.ResponseWriter, req *http.Request) {
	url := "http://localhost:5080/api/default/email"
	r, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Println("No fue posible conectarse ", err)
	} else {
		r.SetBasicAuth(user, pass)
		w.Header().Add("Access-Control-Allow-Origin", "*")
		response, err := client.Do(r)
		if err != nil {
			log.Println("Sin respuesta ", err)
		} else {
			body, err := io.ReadAll(response.Body)
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

	router.Get("/Getemails", GetEmailsData)
	router.Delete("/Delemails", DelEmailsData)
	router.Post("/Postemails", PostEmailsData)

	log.Println(http.ListenAndServe(":3000", router))
}
