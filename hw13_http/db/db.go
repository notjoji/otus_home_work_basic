package db

import (
	"encoding/json"
	"fmt"
)

type Worker struct {
	ID     string
	Name   string
	Salary int
}

type WorkerDB struct {
	Workers map[string]*Worker
}

var WorkerDBInstance *WorkerDB

func NewWorkerDB() *WorkerDB {
	m := map[string]*Worker{
		"1": {
			ID:     "1",
			Name:   "John",
			Salary: 15000,
		},
		"2": {
			ID:     "2",
			Name:   "Mark",
			Salary: 17000,
		},
	}
	WorkerDBInstance = &WorkerDB{m}
	return WorkerDBInstance
}

func Get(id string) ([]byte, error) {
	return json.Marshal(WorkerDBInstance.Workers[id])
}

func Add(data []byte) error {
	var worker *Worker
	err := json.Unmarshal(data, &worker)
	if err != nil {
		return err
	}
	WorkerDBInstance.Workers[worker.ID] = worker
	return nil
}

func Remove(id string) error {
	check := WorkerDBInstance.Workers[id]
	if check == nil {
		return fmt.Errorf("worker not exist with id=%s", id)
	}
	delete(WorkerDBInstance.Workers, id)
	return nil
}

func UpdateWorker(oldWorker, newWorker *Worker) *Worker {
	if newWorker == nil {
		return oldWorker
	}
	if newWorker.Name != "" {
		oldWorker.Name = newWorker.Name
	}
	if newWorker.Salary != 0 {
		oldWorker.Salary = newWorker.Salary
	}
	return oldWorker
}

func Update(id string, data []byte) error {
	check := WorkerDBInstance.Workers[id]
	if check == nil {
		return fmt.Errorf("worker not exist with id=%s", id)
	}

	var worker *Worker
	err := json.Unmarshal(data, &worker)
	if err != nil {
		return err
	}
	WorkerDBInstance.Workers[id] = UpdateWorker(WorkerDBInstance.Workers[id], worker)
	return nil
}

func GetAll() ([]byte, error) {
	workers := make([]*Worker, 0)
	for _, worker := range WorkerDBInstance.Workers {
		workers = append(workers, worker)
	}
	return json.Marshal(workers)
}
