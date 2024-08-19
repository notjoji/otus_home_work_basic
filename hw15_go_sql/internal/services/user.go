package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/internal/repository"
	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/internal/utils"
	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/pkg/pgdb"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
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

	params := repository.GetUsersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	ctx := context.Background()

	tx, err := pgdb.DB.Conn().BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Transaction begin failed"}`))
		return
	}
	repo := repository.New(pgdb.DB.Conn()).WithTx(tx)

	result, err := repo.GetUsers(ctx, params)
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
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

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, []byte(`{"success": false, "msg": "GET method is required"}`))
		return
	}

	userID, err := utils.GetNumericPathParam(r, "/api/getUserByID/", "userID")
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}

	ctx := context.Background()

	tx, err := pgdb.DB.Conn().BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Transaction begin failed"}`))
		return
	}
	repo := repository.New(pgdb.DB.Conn()).WithTx(tx)

	result, err := repo.GetUserById(ctx, int64(userID))
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.ResponseJSON(w, []byte(`{"success": false, "msg": "PUT method is required"}`))
		return
	}

	request, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Cannot read request"}`))
		return
	}

	var params repository.UpdateUserParams
	err = json.Unmarshal(request, &params)
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Cannot unmarshal request"}`))
		return
	}

	ctx := context.Background()

	tx, err := pgdb.DB.Conn().BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Transaction begin failed"}`))
		return
	}
	repo := repository.New(pgdb.DB.Conn()).WithTx(tx)

	_, err = repo.UpdateUser(ctx, params)
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}

	utils.ResponseJSON(w, []byte(
		fmt.Sprintf(`{"success": true,"msg": "User updated, id=%d"}`, params.ID),
	))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ResponseJSON(w, []byte(`{"success": false, "msg": "POST method is required"}`))
		return
	}

	request, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Cannot read request"}`))
		return
	}

	var params repository.CreateUserParams
	err = json.Unmarshal(request, &params)
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Cannot unmarshal request"}`))
		return
	}

	ctx := context.Background()

	tx, err := pgdb.DB.Conn().BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Transaction begin failed"}`))
		return
	}
	repo := repository.New(pgdb.DB.Conn()).WithTx(tx)

	id, err := repo.CreateUser(ctx, params)
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}

	utils.ResponseJSON(w, []byte(
		fmt.Sprintf(`{"success": true,"msg": "User created, id=%d"}`, id),
	))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		utils.ResponseJSON(w, []byte(`{"success": false, "msg": "POST method is required"}`))
		return
	}

	userID, err := utils.GetNumericPathParam(r, "/api/deleteUserByID/", "userID")
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}

	ctx := context.Background()

	tx, err := pgdb.DB.Conn().BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Transaction begin failed"}`))
		return
	}
	repo := repository.New(pgdb.DB.Conn()).WithTx(tx)

	_, err = repo.DeleteUser(ctx, int64(userID))
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}

	utils.ResponseJSON(w, []byte(
		fmt.Sprintf(`{"success": true,"msg": "User created, id=%d"}`, userID),
	))
}
