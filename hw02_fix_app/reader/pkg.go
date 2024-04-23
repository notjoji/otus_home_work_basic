package reader

import (
	"encoding/json"
	"io"
	"os"

	"github.com/notjoji/otus_home_work_basic/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var data []types.Employee

	json.Unmarshal(bytes, &data)

	return data, nil
}
