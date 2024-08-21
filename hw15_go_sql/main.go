package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/internal/config"
	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/internal/services"
	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/pkg/pgdb"
)

func init() {
	if err := config.Load(".env"); err != nil {
		log.Fatal("Didn`t read .env config")
		return
	}
	ctx := context.Background()

	if err := pgdb.New(ctx, os.Getenv("DB_DSN")); err != nil {
		log.Fatal("@[main] can't init service s3client: ", err)
		return
	}
}

func main() {
	port := os.Getenv("API_PORT")

	server := &http.Server{
		Addr:        fmt.Sprintf(":%s", port),
		ReadTimeout: time.Second * 3,
		Handler:     http.HandlerFunc(services.Router),
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(fmt.Errorf("http listen err: %w", err))
	}
}
