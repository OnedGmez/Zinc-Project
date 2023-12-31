package main

import (
	"encoding/json"
	"log"
	"testing"
)

func TestNavDir(t *testing.T) {

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
