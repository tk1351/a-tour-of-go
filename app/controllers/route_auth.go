package controllers

import (
	"a-tour-of-go/app/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("signup")
	} else if r.Method == "POST" {
		length, err := strconv.Atoi(r.Header.Get("Content-Length"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer r.Body.Close()

		body := make([]byte, length)
		length, err = r.Body.Read(body)
		if err != nil && err != io.EOF {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var jsonBody map[string]interface{}
		err = json.Unmarshal(body[:length], &jsonBody)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		user := models.User{
			Name:     jsonBody["name"].(string),
			Email:    jsonBody["email"].(string),
			Password: jsonBody["password"].(string),
		}

		if err := user.CreateUser(); err != nil {
			log.Fatalln(err)
			return
		}
		w.WriteHeader(http.StatusOK)

	}

}
