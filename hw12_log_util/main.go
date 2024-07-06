package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type LogStats struct {
	LogMap map[string]int
}

func LoadLogFileAndCreateLogStats(path string) (LogStats, error) {
	stats := LogStats{LogMap: make(map[string]int)}
	file, err := os.Open(path)
	if err != nil {
		return stats, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		logLevel := strings.Split(line, " ")[0]
		stats.LogMap[logLevel]++
	}
	return stats, nil
}

func WriteLogStatsToFile(path string, stats LogStats, logLevel string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			return
		}
	}(file)

	buf := bufio.NewWriter(file)
	_, err = buf.WriteString(logLevel + ": " + strconv.Itoa(stats.LogMap[logLevel]))
	if err != nil {
		return err
	}
	err = buf.Flush()
	if err != nil {
		return err
	}
	return nil
}

func WriteLogsToStdOut(stats LogStats, logLevel string) error {
	buf := bufio.NewWriter(os.Stdout)
	_, err := buf.WriteString(logLevel + ": " + strconv.Itoa(stats.LogMap[logLevel]))
	if err != nil {
		return err
	}
	err = buf.Flush()
	if err != nil {
		return err
	}
	return nil
}

func GetEnvironmentVars() map[string]string {
	file := flag.String("file", "", "log file path")
	level := flag.String("level", "", "log level")
	output := flag.String("output", "", "output file path")
	flag.Parse()

	if file == nil || *file == "" {
		*file = os.Getenv("LOG_ANALYZER_FILE")
	}
	if level == nil || *level == "" {
		*level = os.Getenv("LOG_ANALYZER_LEVEL")
		if *level == "" {
			*level = "info"
		}
	}
	if output == nil || *output == "" {
		*output = os.Getenv("LOG_ANALYZER_OUTPUT")
	}

	res := make(map[string]string)
	res["file"] = *file
	res["level"] = *level
	res["output"] = *output
	return res
}

func ProcessLogs(vars map[string]string) error {
	if vars["file"] == "" {
		return fmt.Errorf("no log file provided")
	}
	stats, err := LoadLogFileAndCreateLogStats(vars["file"])
	if err != nil {
		return err
	}
	if vars["output"] != "" {
		err = WriteLogStatsToFile(vars["output"], stats, vars["level"])
		if err != nil {
			return err
		}
	} else {
		err = WriteLogsToStdOut(stats, vars["level"])
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	err := ProcessLogs(GetEnvironmentVars())
	if err != nil {
		log.Fatal(err)
	}
}
