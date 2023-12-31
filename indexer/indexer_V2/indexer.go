package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/mail"
	"os"
	"regexp"
	"strings"
	"sync"
)

const (
	principalPath   = "../enron_mail_20110402/maildir"
	dataIndexedPath = "data.ndjson"
	user            = "ogsystem@gmail.com"
	pass            = "Complexpass#123"
	url             = "http://localhost:5080/api/default/email/_multi"
	urlEP           = "http://localhost:3000/Postemails_V2"
)

type emailSt struct {
	Id      string
	Date    string
	From    string
	To      string
	Subject string
	Body    string
}

var (
	//CHANNELS
	emailsPath = make(chan string)
	emailsData = make(chan emailSt)

	//CONTROL VARIABLES
	jsonData string
	i        int = 0
	wge      sync.WaitGroup
	r        = regexp.MustCompile(`\t+`)
	n        = regexp.MustCompile(`\n+`)
	client   = http.Client{}

	//DATA CHANGING
	to      string
	from    string
	date    string
	subject string
)

func main() {

	go func() {
		for p := range emailsPath {
			processEmail(p, &wge)
		}
	}()

	go func() {
		for d := range emailsData {
			wge.Add(1)
			i++
			if d.Id != "" {
				j, err := json.Marshal(d)
				if err != nil {
					log.Println("No se pudo convertir a ndJson", err)
				} else {
					jsonData = concatString(jsonData, string(j), "\n")
					if i >= 15000 {
						createIndex(jsonData)
						jsonData = ""
					}
				}
			}
			wge.Done()
		}
	}()

	navDir(principalPath, &wge)
	if jsonData != "" {
		createIndex(jsonData)
		jsonData = ""
	}
	wge.Wait()
	close(emailsData)
}

/*
navDir es la función utilizada para navegar por los directorios y obtener las rutas de acceso a los archivos que contiene los emails
@param path: string, contiene la ruta de acceso a los directorios y/o los archivos
@param wge: sync.WaitGroup, se utiliza para controlar las Go rutines
*/
func navDir(path string, wge *sync.WaitGroup) {
	dir, err := os.ReadDir(path)
	if err != nil {
		log.Println("No se puede leer la dirección: ", err)
	}
	for _, file := range dir {
		if file.Type().IsDir() {
			navDir(path+"/"+file.Name(), wge)
		} else {
			wge.Add(1) //Agregamos 1 unidad por cada ruta de archivo de email encontrado
			emailsPath <- path + "/" + file.Name()
		}
	}
}

/*
processEmail es la función utilizada para extraer la información de cada archivo y enviarlo al channel de estructura
@param path: string, contiene la ruta del archivo del que se extraerá la información
@param wge: sync.WaitGroup, se utiliza para controlar las Go rutines
*/
func processEmail(path string, wge *sync.WaitGroup) {
	defer wge.Done()
	content, err := os.ReadFile(path)
	if err != nil {
		log.Println("Error al intentar abrir el archivo", err)
	}

	//Reformatear el contenido del archivo
	r := strings.NewReader(string(content))
	msg, err := mail.ReadMessage(r)
	if err == nil {
		header := msg.Header
		to = oneLine(header.Get("To"), "\n")
		from = oneLine(header.Get("From"), "\n")
		date = header.Get("Date")
		subject = header.Get("Subject")
		body, err := io.ReadAll(msg.Body)
		if err != nil {
			log.Println("No se pudo extraer el conenido ", err)
		} else {
			emailsData <- emailSt{
				Id:      path,
				Date:    date,
				From:    strings.ToLower(from),
				To:      strings.ToLower(to),
				Subject: subject,
				Body:    formatContent(string(body)),
			}
		}
	} else {
		log.Println("No se pudo leer el contenido del correo: ", path, " ", err)
	}
}

/*
formatContent es la función utilizada para agregarle el signo de scape(\) a los signos como ",\ para que permanezcan en el string del mensaje, así como reemplazar las sangrías por una marca que indique donde existía
@param content: string, contiene el string que se requiere darle el formato
@return string con el contenido ya formateado
*/
func formatContent(content string) string {
	var returnData string
	if content != "" {
		content = strings.TrimSpace(content)
		content = strings.ReplaceAll(content, `\`, ` \\ `)
		content = (string(r.ReplaceAll([]byte(content), []byte(" [.SANGRIA.] "))))
		content = (string(n.ReplaceAll([]byte(content), []byte(" [.SALTO.] "))))
		content = strings.ReplaceAll(content, `"`, `\"`)
		returnData = concatString(returnData, content, " ")
	}
	return returnData
}

/*
oneLine es la función utilizada para convertir múltiples líneas en una sola línea de string
@param content: string, contiene el string que se requiere darle el formato
@return string con el contenido ya formateado
*/
func oneLine(content string, sep string) string {
	var returnData string
	dataTMP := strings.Split(string(content), sep)
	for _, text := range dataTMP {
		if text != "" {
			text = strings.TrimSpace(text)
			returnData = concatString(returnData, text, " ")
		}
	}
	return returnData
}

func concatString(variable string, data string, joinS string) string {
	if variable != "" {
		variable = variable + joinS + data
	} else {
		variable = data
	}
	return variable
}

/*
createIndex es la función utilizada para enviar la data a OpenObserve
@param content: string, contiene el string con la data a enviar a Openobserve
*/
func createIndex(content string) {
	req, err := http.NewRequest("POST", url, strings.NewReader(content))
	if err != nil {
		log.Println("No fue posible conectarse ", err)
	} else {
		req.SetBasicAuth(user, pass)
		r, err := client.Do(req)
		if err != nil {
			log.Println("Sin respuesta ", err)
		} else {
			defer r.Body.Close()
		}
	}
}

// /*
// RETIRADO POR CONSUMO EXCESIVO EN RENDIMIENTO
// createIndex es la función utilizada para enviar la data a OpenObserve
// @param content: string, contiene el string con la data a enviar a Openobserve
// */
// func createIndex(content string) {
// 	req, err := http.NewRequest("POST", urlEP, strings.NewReader(content))
// 	if err != nil {
// 		log.Println("No fue posible conectarse ", err)
// 	} else {
// 		req.SetBasicAuth(user, pass)
// 		r, err := client.Do(req)
// 		if err != nil {
// 			log.Println("Sin respuesta ", err)
// 		} else {
// 			defer r.Body.Close()
// 		}
// 	}
// }
