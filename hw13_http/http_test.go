package main

import (
	"testing"
	"time"

	Client "github.com/notjoji/otus_home_work_basic/hw13_http/client"
	DB "github.com/notjoji/otus_home_work_basic/hw13_http/db"
	Server "github.com/notjoji/otus_home_work_basic/hw13_http/server"

	"github.com/stretchr/testify/assert"
)

func TestHttp(t *testing.T) {
	_ = DB.NewWorkerDB()
	go Server.ListenAndServe()
	to := time.After(time.Second)
	<-to
	assert.NoError(t, Client.Start())
}
