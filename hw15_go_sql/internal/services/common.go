package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/internal/repository"
	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/internal/utils"
	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/pkg/pgdb"
)

type PageableMethod string

const (
	UsersAndProducts PageableMethod = "getUsersAndProducts"
	Users            PageableMethod = "getUsers"
)

func GetPageableEntities(w http.ResponseWriter, r *http.Request, method PageableMethod) {
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, []byte(`{"success": false, "msg": "GET method is required"}`))
		return
	}

	limit, offset, err := utils.GetPageableParams(r)
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}

	ctx := context.Background()

	repo := repository.New(pgdb.DB.Conn())

	var result any
	var methodError error
	switch method {
	case UsersAndProducts:
		result, methodError = repo.GetUsersAndProducts(ctx, repository.GetUsersAndProductsParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	case Users:
		result, methodError = repo.GetUsers(ctx, repository.GetUsersParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	}
	if methodError != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, methodError.Error()),
		))
		return
	}

	resultByte, err := json.Marshal(result)
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}
	utils.ResponseJSON(w, resultByte)
}
