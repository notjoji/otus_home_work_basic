package main

import (
	"github.com/notjoji/otus_home_work_basic/hw02_fix_app/reader"
	"github.com/notjoji/otus_home_work_basic/hw02_fix_app/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadJSON(t *testing.T) {
	testCases := []struct {
		name      string
		filePath  string
		error     string
		employees []types.Employee
	}{
		{
			"incorrect_path",
			"data1.json",
			"open data1.json: The system cannot find the file specified.",
			nil,
		},
		{
			"unmarshal_error",
			"unmarshal_error_data.json",
			"json: cannot unmarshal number 9223372036854775808 into Go struct field Employee.departmentId of type int",
			nil,
		},
		{
			"positive_case",
			"data.json",
			"",
			[]types.Employee{
				{
					UserID:       10,
					Age:          25,
					Name:         "Rob",
					DepartmentID: 3,
				},
				{
					UserID:       11,
					Age:          30,
					Name:         "George",
					DepartmentID: 2,
				},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expected, err := reader.ReadJSON(testCase.filePath)
			if err != nil {
				assert.Equal(t, testCase.error, err.Error())
			}
			if expected != nil {
				for index, employee := range testCase.employees {
					assert.Equal(t, employee, expected[index])
				}
			}
		})
	}
}
