package types

import "fmt"

type Employee struct {
	UserID       int    `json:"userId"`
	Age          int    `json:"age"`
	Name         string `json:"name"`
	DepartmentID int    `json:"departmentId"`
}

func (e Employee) String() string {
	return fmt.Sprintf("%d %d %s %d", e.UserID, e.Age, e.Name, e.DepartmentID)
}
