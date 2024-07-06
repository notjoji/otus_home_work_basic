package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ReadOutFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			return
		}
	}(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func TestProcessLogs(t *testing.T) {
	testCases := []struct {
		name   string
		file   string
		level  string
		output string
		want   string
	}{
		{
			name:   "positive scenario, default level (stdout)",
			file:   "log1.txt",
			level:  "info",
			output: "",
			want:   "",
		},
		{
			name:   "positive scenario, default level (file)",
			file:   "log1.txt",
			level:  "info",
			output: "out.txt",
			want:   "info: 5",
		},
		{
			name:   "no vars - file path required error",
			file:   "",
			level:  "",
			output: "",
			want:   "",
		},
		{
			name:   "different existing log level",
			file:   "log1.txt",
			level:  "warn",
			output: "out.txt",
			want:   "warn: 2",
		},
		{
			name:   "non-existing log level",
			file:   "log1.txt",
			level:  "trace",
			output: "out.txt",
			want:   "trace: 0",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			vars := map[string]string{"file": tc.file, "level": tc.level, "output": tc.output}
			err := ProcessLogs(vars)
			if vars["file"] == "" {
				assert.Errorf(t, err, "no log file provided")
			} else if vars["output"] != "" {
				actual := ReadOutFile(tc.output)
				assert.Equal(t, tc.want, actual)
			}
		})
	}
}
