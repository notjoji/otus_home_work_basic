package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	DB "github.com/notjoji/otus_home_work_basic/hw13_http/db"
)

func GetAll(client http.Client) error {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet,
		"http://localhost:8080/workers/", nil)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body error: %w", err)
	}
	fmt.Println("GET ALL WORKERS response:", string(result))
	return nil
}

func GetByID(client http.Client, id string) error {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet,
		fmt.Sprintf("http://localhost:8080/workers/%s", id), nil)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body error: %w", err)
	}
	fmt.Println("GET WORKER BY ID response:", string(result))
	return nil
}

func AddNewWorker(client http.Client, worker DB.Worker) error {
	marshalled, err := json.Marshal(worker)
	if err != nil {
		return fmt.Errorf("json marshal error: %w", err)
	}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost,
		"http://localhost:8080/workers/", bytes.NewBuffer(marshalled))
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body error: %w", err)
	}
	fmt.Println("POST NEW WORKER response:", string(result))
	return nil
}

func UpdateWorker(client http.Client, worker DB.Worker, id string) error {
	marshalled, err := json.Marshal(worker)
	if err != nil {
		return fmt.Errorf("json marshal error: %w", err)
	}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPut,
		fmt.Sprintf("http://localhost:8080/workers/%s", id), bytes.NewBuffer(marshalled))
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body error: %w", err)
	}
	fmt.Println("PUT UPDATE WORKER response:", string(result))
	return nil
}

func DeleteWorker(client http.Client, id string) error {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodDelete,
		fmt.Sprintf("http://localhost:8080/workers/%s", id), nil)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body error: %w", err)
	}
	fmt.Println("DELETE WORKER response:", string(result))
	return nil
}

func Start() error {
	client := http.Client{}
	err := GetAll(client)
	if err != nil {
		return err
	}
	err = GetByID(client, "2")
	if err != nil {
		return err
	}
	newWorker := DB.Worker{
		ID:     "3",
		Name:   "Viktor",
		Salary: 13000,
	}
	err = AddNewWorker(client, newWorker)
	if err != nil {
		return err
	}
	updateWorker := DB.Worker{
		Salary: 17500,
	}
	err = UpdateWorker(client, updateWorker, "1")
	if err != nil {
		return err
	}
	err = GetAll(client)
	if err != nil {
		return err
	}
	err = DeleteWorker(client, "2")
	if err != nil {
		return err
	}
	err = GetAll(client)
	if err != nil {
		return err
	}
	return nil
}
