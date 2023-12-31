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
			if d.Id != "" {
				j, err := json.Marshal(d)
				if err != nil {
					log.Println("No se pudo convertir a ndJson", err)
				} else {
					createIndex(j)
				}
			}
			wge.Done()
		}
	}()

	navDir(principalPath, &wge)
	wge.Wait()
	close(emailsData)
}
