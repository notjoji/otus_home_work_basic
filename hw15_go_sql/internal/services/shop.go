package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/internal/repository"
	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/internal/utils"
	"github.com/notjoji/otus_home_work_basic/hw15_go_sql/pkg/pgdb"
)

func GetOrdersByUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userIDParam := r.URL.Query().Get("userID")
	if userIDParam == "" {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Query parameter 'userID' is required"}`))
		return
	}
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Parameter 'userID' is invalid"}`))
		return
	}

	tx, err := pgdb.DB.Conn().BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Transaction begin failed"}`))
		return
	}
	repo := repository.New(pgdb.DB.Conn()).WithTx(tx)

	result, err := repo.GetOrdersByUserId(ctx, int64(userID))
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}

	resulByte, err := json.Marshal(result)
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}
	utils.ResponseJSON(w, resulByte)
}

func GetUsersAndProducts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	limitParam := r.URL.Query().Get("limit")
	offsetParam := r.URL.Query().Get("offset")

	limit := int64(5)
	offset := int64(0)
	var err error
	if limitParam != "" {
		limit, err = strconv.ParseInt(limitParam, 10, 64)
		if err != nil {
			utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Parameter 'limit' is invalid"}`))
			return
		}
	}
	if offsetParam != "" {
		offset, err = strconv.ParseInt(offsetParam, 10, 64)
		if err != nil {
			utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Parameter 'offset' is invalid"}`))
			return
		}
	}

	params := repository.GetUsersAndProductsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	tx, err := pgdb.DB.Conn().BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Transaction begin failed"}`))
		return
	}
	repo := repository.New(pgdb.DB.Conn()).WithTx(tx)

	result, err := repo.GetUsersAndProducts(ctx, params)
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}

	resulByte, err := json.Marshal(result)
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}
	utils.ResponseJSON(w, resulByte)
}

func GetUserStatistics(w http.ResponseWriter, _ *http.Request) {
	ctx := context.Background()

	tx, err := pgdb.DB.Conn().BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		utils.ResponseJSON(w, []byte(`{"success": false,"msg": "Transaction begin failed"}`))
		return
	}
	repo := repository.New(pgdb.DB.Conn()).WithTx(tx)

	result, err := repo.GetUserStatistics(ctx)
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}

	resulByte, err := json.Marshal(result)
	if err != nil {
		utils.ResponseJSON(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}
	utils.ResponseJSON(w, resulByte)
}
