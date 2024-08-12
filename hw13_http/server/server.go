package server

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	Database "github.com/notjoji/otus_home_work_basic/hw13_http/db"
)

func WorkersHandler(w http.ResponseWriter, r *http.Request) {
	var response, request []byte
	var err error
	switch r.Method {
	case "GET":
		id := strings.TrimPrefix(r.URL.Path, "/workers/")
		if id != "" {
			response, err = Database.Get(id)
		} else {
			response, err = Database.GetAll()
		}
	case "POST", "PUT":
		request, err = io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(err.Error()))
		}
		if r.Method == "POST" {
			err = Database.Add(request)
			response = []byte("Сотрудник добавлен")
		} else {
			id := strings.TrimPrefix(r.URL.Path, "/workers/")
			if id != "" {
				err = Database.Update(id, request)
				response = []byte("Сотрудник изменен")
			} else {
				w.WriteHeader(http.StatusBadRequest)
				_, err = w.Write([]byte("Требуется идентификатор сотрудника!"))
			}
		}
	case "DELETE":
		id := strings.TrimPrefix(r.URL.Path, "/workers/")
		if id != "" {
			err = Database.Remove(id)
			response = []byte("Сотрудник удален")
		} else {
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte("Требуется идентификатор сотрудника!"))
		}
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		return
	}

	if response != nil {
		_, err = w.Write(response)
	}
}

func ListenAndServe() {
	http.HandleFunc("/workers/", WorkersHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(fmt.Errorf("http listen err: %v", err))
	}
}
