package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	DB "github.com/notjoji/otus_home_work_basic/hw13_http/db"
)

func GetAll() error {
	resp, err := http.Get("http://localhost:8080/workers")
	if err != nil {
		return fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body error: %v", err)
	}
	fmt.Println("GET ALL WORKERS response:", string(result))
	return nil
}

func GetById(id string) error {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/workers/%s", id))
	if err != nil {
		return fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body error: %v", err)
	}
	fmt.Println("GET WORKER BY ID response:", string(result))
	return nil
}

func AddNewWorker(worker DB.Worker) error {
	marshalled, err := json.Marshal(worker)
	if err != nil {
		return fmt.Errorf("json marshal error: %v", err)
	}
	resp, err := http.Post("http://localhost:8080/workers/", "", bytes.NewBuffer(marshalled))
	if err != nil {
		return fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body error: %v", err)
	}
	fmt.Println("POST NEW WORKER response:", string(result))
	return nil
}

func UpdateWorker(client http.Client, worker DB.Worker, id string) error {
	marshalled, err := json.Marshal(worker)
	if err != nil {
		return fmt.Errorf("json marshal error: %v", err)
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/workers/%s", id),
		bytes.NewBuffer(marshalled))
	if err != nil {
		return fmt.Errorf("request error: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body error: %v", err)
	}
	fmt.Println("PUT UPDATE WORKER response:", string(result))
	return nil
}

func DeleteWorker(client http.Client, id string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/workers/%s", id), nil)
	if err != nil {
		return fmt.Errorf("request error: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body error: %v", err)
	}
	fmt.Println("DELETE WORKER response:", string(result))
	return nil
}

func Start() error {
	err := GetAll()
	if err != nil {
		return err
	}
	err = GetById("2")
	if err != nil {
		return err
	}
	newWorker := DB.Worker{
		ID:     "3",
		Name:   "Viktor",
		Salary: 13000,
	}
	err = AddNewWorker(newWorker)
	if err != nil {
		return err
	}
	client := http.Client{}
	updateWorker := DB.Worker{
		Salary: 17500,
	}
	err = UpdateWorker(client, updateWorker, "1")
	if err != nil {
		return err
	}
	err = GetAll()
	if err != nil {
		return err
	}
	err = DeleteWorker(client, "2")
	if err != nil {
		return err
	}
	err = GetAll()
	if err != nil {
		return err
	}
	return nil
}
